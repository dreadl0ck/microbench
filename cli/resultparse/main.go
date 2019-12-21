package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type job struct {
	fileKeywords        []string
	excludeFileKeywords []string
	lineIdent           string
	lineRegex           *regexp.Regexp
	name                string
}

var defaultRegEx = regexp.MustCompile("delta=[0-9]*.[0-9]*m?s")
var wg sync.WaitGroup

const (
	identKernelBootTime = "kernel boot time received"
	identHashBenchmark = "hash loop benchmark"
	identWebService = "time until HTTP reply from webservice"
	identShutdownTime = "shutdown complete"
	identKernelLogLines = "number of kernel log lines received"
)

func main() {

	sequentialJobs := []job{
		{
			fileKeywords:        []string{"qemu", "sequential"},
			excludeFileKeywords: []string{"emulated"},
			name:                "host_cpu",
		},
		{
			fileKeywords: 	[]string{"qemu", "sequential", "emulated"},
			name: 			"emulated_cpu",
		},
		{
			fileKeywords: 	[]string{"firecracker", "sequential"},
			name: 			"T2_cpu",
		},
		{
			fileKeywords: 	[]string{"firecracker", "sequential", "C3"},
			name: 			"C3_cpu",
		},
		{
			fileKeywords: 	[]string{"firecracker", "sequential", "default", "kernel"},
			name: 			"default_kernel",
		},
	}

	concurrentJobs := []job{
		{
			fileKeywords:        []string{"qemu", "x10"},
			excludeFileKeywords: []string{"emulated"},
		},
		{
			fileKeywords:        []string{"qemu", "x10", "emulated"},
		},
		{
			fileKeywords:        []string{"qemu", "x20"},
			excludeFileKeywords: []string{"emulated"},
		},
		{
			fileKeywords:        []string{"qemu", "x20", "emulated"},
		},
		{
			fileKeywords:        []string{"firecracker", "x10"},
		},
		{
			fileKeywords:        []string{"firecracker", "x20"},
		},
	}

	// mean-hashing-time-sequential.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-hashing-time-sequential.py",
		lineIdent: identHashBenchmark,
		jobs: sequentialJobs,
	})

	// mean-hashing-time-concurrent.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-hashing-time-concurrent.py",
		lineIdent: identHashBenchmark,
		jobs:      concurrentJobs,
	})

	// mean-kernel-boot-time-sequential.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-kernel-boot-time-sequential.py",
		lineIdent: identKernelBootTime,
		jobs: sequentialJobs,
	})

	// mean-kernel-boot-time-concurrent.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-kernel-boot-time-concurrent.py",
		lineIdent: identKernelBootTime,
		jobs: concurrentJobs,
	})

	// mean-webservice-time-sequential.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-webservice-time-sequential.py",
		lineIdent: identWebService,
		jobs: sequentialJobs,
	})

	// mean-webservice-time-concurrent.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-webservice-time-concurrent.py",
		lineIdent: identWebService,
		jobs: concurrentJobs,
	})

	// mean-shutdown-time-sequential.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-shutdown-time-sequential.py",
		lineIdent: identShutdownTime,
		jobs: sequentialJobs,
	})

	// mean-shutdown-time-concurrent.png
	go generate(plot{
		mean: true,
		script: "plots/scripts/mean-shutdown-time-concurrent.py",
		lineIdent: identShutdownTime,
		jobs: concurrentJobs,
	})

	// webservice-time-sequential.png
	go generate(plot{
		script: "plots/scripts/webservice-time-sequential.py",
		lineIdent: identWebService,
		jobs: sequentialJobs,
		forceName: true,
	})

	// webservice-time-concurrent.png
	go generate(plot{
		script: "plots/scripts/webservice-time-concurrent.py",
		lineIdent: identWebService,
		jobs: concurrentJobs,
	})

	// kernel-boot-time-sequential.png
	go generate(plot{
		script: "plots/scripts/kernel-boot-time-sequential.py",
		lineIdent: identKernelBootTime,
		jobs: sequentialJobs,
		forceName: true,
	})

	// webservice-time-concurrent.png
	go generate(plot{
		script: "plots/scripts/kernel-boot-time-concurrent.py",
		lineIdent: identKernelBootTime,
		jobs: concurrentJobs,
	})

	// shutdown-time-sequential.png
	go generate(plot{
		script:    "plots/scripts/shutdown-time-sequential.py",
		lineIdent: identShutdownTime,
		jobs:      sequentialJobs,
		forceName: true,
	})

	// shutdown-time-concurrent.png
	go generate(plot{
		script: "plots/scripts/shutdown-time-concurrent.py",
		lineIdent: identShutdownTime,
		jobs: concurrentJobs,
	})

	// kernel-log-entries.png
	go generate(plot{
		mean: true,
		reg: regexp.MustCompile("lines=[0-9]*"),
		script: "plots/scripts/kernel-log-entries.py",
		lineIdent: identKernelLogLines,
		jobs: sequentialJobs,
	})

	time.Sleep(500 * time.Millisecond)
	wg.Wait()
}

type plot struct {
	mean bool
	reg *regexp.Regexp
	forceName bool
	script string
	lineIdent string
	jobs []job
}

//func generate(mean bool, reg *regexp.Regexp, script string, lineIdent string, jobs ...*job) {
func generate(p plot) {

	wg.Add(1)
	defer wg.Done()

	var copiedJobs []job

	for _, j := range p.jobs {
		var n = j.name
		if p.forceName || j.name == "" {
			n = strings.Join(j.fileKeywords, "_")
		}

		// make a deep copy of the jobs
		// because the name modification by the force flag can happen in parallel
		copiedJobs = append(copiedJobs, job{
			lineIdent: j.lineIdent,
			lineRegex: j.lineRegex,
			name: n,
			fileKeywords: j.fileKeywords,
			excludeFileKeywords: j.excludeFileKeywords,
		})
	}

	var data = make(map[string][]float64)

	files, err := ioutil.ReadDir("experiment_logs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		parseJobs(data, f, p.reg, p.lineIdent, copiedJobs)
	}

	var out string

	for key, values := range data {
		out += key + " = [\n"
		lastElem := len(values) - 1
		for i, v := range values {
			if i == lastElem {
				out += "\t" + strconv.FormatFloat(v, 'f', 5, 64) + "\n"
			} else {
				out += "\t" + strconv.FormatFloat(v, 'f', 5, 64) + ",\n"
			}
		}
		out += "]\n\n"
	}

	//fmt.Println(out)

	t, err := template.ParseFiles(p.script)
	if err != nil {
		log.Fatal(err)
	}

	outScriptPath := filepath.Dir(p.script) + "/generated/" + strings.TrimSuffix(filepath.Base(p.script), ".py") + ".py"
	f, err := os.Create(outScriptPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var logStatements string
	for k := range data {
		logStatements += "print(stats.mean("+ k + "))\n"
	}

	var (
		current int
		load string
		objects string
		lastElem = len(data)
	)
	for _, j := range copiedJobs {
		if _, ok := data[j.name]; ok {

			current++
			if current == lastElem {
				objects += "'" + strings.ReplaceAll(j.name, "_", " ") + "'"
				if p.mean {
					load += "\tstats.mean("+ j.name + ")"
				} else {
					load += "\t"+ j.name + ""
				}
			} else {
				objects += "'" + strings.ReplaceAll(j.name, "_", " ") + "',"
				if p.mean {
					load += "\tstats.mean("+ j.name + "),\n"
				} else {
					load += "\t"+ j.name + ",\n"
				}
			}
		} else {
			fmt.Println(p.script + ": no data for", j.name, "found after collecting values")
		}
	}

	outImageName := filepath.Dir(p.script) + "/images/" + strings.TrimSuffix(filepath.Base(p.script), ".py") + ".png"

	err = t.Execute(f, struct{
		Data string
		Log  string
		Load string
		Objects string
		Out string
	}{
		out,
		logStatements,
		load,
		objects,
		"'"+ outImageName + "'",
	})
	if err != nil {
		log.Fatal(err)
	}

	outCmd, err := exec.Command("python3", outScriptPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(outCmd))
		log.Fatal(err)
	}
	//fmt.Println(string(outCmd))

	fmt.Println("created", outImageName)
}

func parseJobs(data map[string][]float64, f os.FileInfo, reg *regexp.Regexp, lineIdent string, jobs []job) {

	if reg == nil {
		reg = defaultRegEx
	}

	for _, j := range jobs {

		var nextJob bool
		for _, w := range j.fileKeywords {
			if !strings.Contains(f.Name(), w) {
				nextJob = true
				break
			}
		}

		for _, w := range j.excludeFileKeywords {
			if strings.Contains(f.Name(), w) {
				nextJob = true
				break
			}
		}

		if nextJob {
			continue
		}

		contents, err := ioutil.ReadFile("experiment_logs/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, l := range strings.Split(string(contents), "\n") {

			if !strings.Contains(l, lineIdent) {
				continue
			}

			value := reg.FindString(l)

			// remove the XXX= prefix from the extracted value
			for i, c := range value {
				if c == '=' {
					value = value[i+1:]
				}
			}

			//fmt.Println(f.Name(), value, l)

			var toMs bool
			if !strings.HasSuffix(value, "ms") && strings.HasSuffix(value, "s") {
				toMs = true
			}
			final := strings.TrimSuffix(
				strings.TrimSuffix(value, "ms"),
			"s")

			if final == "" {
				continue
			}

			f, err := strconv.ParseFloat(
				final,
				64)
			if err != nil {
				log.Fatal("failed to parse float ", final, ", error: ", err)
			}

			if toMs {
				f = f * 1000
			}

			data[j.name] = append(data[j.name], f)

			// next iteration
			continue
		}
	}
}
