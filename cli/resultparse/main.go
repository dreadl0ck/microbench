/*
 * MICROBENCH - A testbed for comparing microvm technologies
 * Copyright (c) 2019 Philipp Mieden and Philippe Partarrieu
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/mgutz/ansi"
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
	name                string
	ident               string
}

var (
	flagDebug    = flag.Bool("debug", false, "toggle debug mode")
	defaultRegEx = regexp.MustCompile("delta=[0-9]*.[0-9]*m?s")
	wg           sync.WaitGroup

	identKernelBootTime = "kernel boot time received"
	identHashBenchmark  = "hash loop benchmark"
	identWebService     = "time until HTTP reply from webservice"
	identShutdownTime   = "shutdown complete"
	identKernelLogLines = "number of kernel log lines received"
)

func main() {

	flag.Parse()

	sequentialJobs := []job{
		{
			fileKeywords:        []string{"qemu", "sequential"},
			excludeFileKeywords: []string{"emulated"},
			name:                "host_cpu",
		},
		{
			fileKeywords: []string{"qemu", "sequential", "emulated"},
			name:         "emulated_cpu",
		},
		{
			fileKeywords: []string{"firecracker", "sequential"},
			name:         "T2_cpu",
		},
		{
			fileKeywords: []string{"firecracker", "sequential", "C3"},
			name:         "C3_cpu",
		},
		{
			fileKeywords: []string{"firecracker", "sequential", "default", "kernel"},
			name:         "default_kernel",
		},
	}

	concurrentJobs := []job{
		{
			fileKeywords:        []string{"qemu", "x10"},
			excludeFileKeywords: []string{"emulated"},
		},
		{
			fileKeywords: []string{"qemu", "x10", "emulated"},
		},
		{
			fileKeywords:        []string{"qemu", "x20"},
			excludeFileKeywords: []string{"emulated"},
		},
		{
			fileKeywords: []string{"qemu", "x20", "emulated"},
		},
		{
			fileKeywords: []string{"firecracker", "x10"},
		},
		{
			fileKeywords: []string{"firecracker", "x20"},
		},
	}

	// mean-hashing-time-sequential.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-hashing-time-sequential.py",
		lineIdent: identHashBenchmark,
		jobs:      sequentialJobs,
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
		mean:      true,
		script:    "plots/scripts/mean-kernel-boot-time-sequential.py",
		lineIdent: identKernelBootTime,
		jobs:      sequentialJobs,
	})

	// mean-kernel-boot-time-concurrent.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-kernel-boot-time-concurrent.py",
		lineIdent: identKernelBootTime,
		jobs:      concurrentJobs,
	})

	// mean-webservice-time-sequential.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-webservice-time-sequential.py",
		lineIdent: identWebService,
		jobs:      sequentialJobs,
	})

	// mean-webservice-time-concurrent.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-webservice-time-concurrent.py",
		lineIdent: identWebService,
		jobs:      concurrentJobs,
	})

	// mean-shutdown-time-sequential.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-shutdown-time-sequential.py",
		lineIdent: identShutdownTime,
		jobs:      sequentialJobs,
	})

	// mean-shutdown-time-concurrent.png
	go generate(plot{
		mean:      true,
		script:    "plots/scripts/mean-shutdown-time-concurrent.py",
		lineIdent: identShutdownTime,
		jobs:      concurrentJobs,
	})

	// webservice-time-sequential.png
	go generate(plot{
		script:    "plots/scripts/webservice-time-sequential.py",
		lineIdent: identWebService,
		jobs:      sequentialJobs,
		forceName: true,
	})

	// webservice-time-concurrent.png
	go generate(plot{
		script:    "plots/scripts/webservice-time-concurrent.py",
		lineIdent: identWebService,
		jobs:      concurrentJobs,
	})

	// kernel-boot-time-sequential.png
	go generate(plot{
		script:    "plots/scripts/kernel-boot-time-sequential.py",
		lineIdent: identKernelBootTime,
		jobs:      sequentialJobs,
		forceName: true,
	})

	// webservice-time-concurrent.png
	go generate(plot{
		script:    "plots/scripts/kernel-boot-time-concurrent.py",
		lineIdent: identKernelBootTime,
		jobs:      concurrentJobs,
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
		script:    "plots/scripts/shutdown-time-concurrent.py",
		lineIdent: identShutdownTime,
		jobs:      concurrentJobs,
	})

	// kernel-log-entries.png
	go generate(plot{
		mean:      true,
		reg:       regexp.MustCompile("lines=[0-9]*"),
		script:    "plots/scripts/kernel-log-entries.py",
		lineIdent: identKernelLogLines,
		jobs:      sequentialJobs,
	})

	// kernel-boot-and-webservice-time-concurrent.png
	go generate(plot{
		mean:      true,
		stacked:   true,
		script:    "plots/scripts/kernel-boot-and-webservice-time-concurrent.py",
		lineIdent: identKernelBootTime,
		jobs: []job{

			// kernel boot values (default)
			{
				fileKeywords:        []string{"qemu", "x10"},
				excludeFileKeywords: []string{"emulated"},
				name:                "qemu_x10_kernel_boot",
			},
			{
				fileKeywords: []string{"qemu", "x10", "emulated"},
				name:         "qemu_x10_emulated_kernel_boot",
			},
			{
				fileKeywords:        []string{"qemu", "x20"},
				excludeFileKeywords: []string{"emulated"},
				name:                "qemu_x20_kernel_boot",
			},
			{
				fileKeywords: []string{"qemu", "x20", "emulated"},
				name:         "qemu_x20_emulated_kernel_boot",
			},
			{
				fileKeywords: []string{"firecracker", "x10"},
				name:         "firecracker_x10_kernel_boot",
			},
			{
				fileKeywords: []string{"firecracker", "x20"},
				name:         "firecracker_x20_kernel_boot",
			},

			// webservice values
			{
				fileKeywords:        []string{"qemu", "x10"},
				excludeFileKeywords: []string{"emulated"},
				name:                "qemu_x10_webservice",
				ident:               identWebService,
			},
			{
				fileKeywords: []string{"qemu", "x10", "emulated"},
				name:         "qemu_x10_webservice",
				ident:        identWebService,
			},
			{
				fileKeywords:        []string{"qemu", "x20"},
				excludeFileKeywords: []string{"emulated"},
				name:                "qemu_x20_webservice",
				ident:               identWebService,
			},
			{
				fileKeywords: []string{"qemu", "x20", "emulated"},
				name:         "qemu_x20_emulated_webservice",
				ident:        identWebService,
			},
			{
				fileKeywords: []string{"firecracker", "x10"},
				name:         "firecracker_x10_webservice",
				ident:        identWebService,
			},
			{
				fileKeywords: []string{"firecracker", "x20"},
				name:         "firecracker_x20_webservice",
				ident:        identWebService,
			},
		},
	})

	// kernel-boot-and-webservice-time-sequential.png
	go generate(plot{
		mean:      true,
		stacked:   true,
		script:    "plots/scripts/kernel-boot-and-webservice-time-sequential.py",
		lineIdent: identKernelBootTime,
		jobs: []job{

			// kernel boot values (default)
			{
				fileKeywords:        []string{"qemu", "sequential"},
				excludeFileKeywords: []string{"emulated"},
				name:                "host_cpu_kernel_boot",
			},
			{
				fileKeywords: []string{"qemu", "sequential", "emulated"},
				name:         "emulated_cpu_kernel_boot",
			},
			{
				fileKeywords: []string{"firecracker", "sequential"},
				name:         "T2_cpu_kernel_boot",
			},
			{
				fileKeywords: []string{"firecracker", "sequential", "C3"},
				name:         "C3_cpu_kernel_boot",
			},
			{
				fileKeywords: []string{"firecracker", "sequential", "default", "kernel"},
				name:         "default_kernel_kernel_boot",
			},

			// webservice values
			{
				fileKeywords:        []string{"qemu", "sequential"},
				excludeFileKeywords: []string{"emulated"},
				name:                "host_cpu_webservice",
				ident:               identWebService,
			},
			{
				fileKeywords: []string{"qemu", "sequential", "emulated"},
				name:         "emulated_cpu_webservice",
				ident:        identWebService,
			},
			{
				fileKeywords: []string{"firecracker", "sequential"},
				name:         "T2_cpu_webservice",
				ident:        identWebService,
			},
			{
				fileKeywords: []string{"firecracker", "sequential", "C3"},
				name:         "C3_cpu_webservice",
				ident:        identWebService,
			},
			{
				fileKeywords: []string{"firecracker", "sequential", "default", "kernel"},
				name:         "default_kernel_webservice",
				ident:        identWebService,
			},
		},
	})

	time.Sleep(500 * time.Millisecond)
	wg.Wait()
}

type plot struct {
	mean      bool
	stacked   bool
	reg       *regexp.Regexp
	forceName bool
	script    string
	lineIdent string
	jobs      []job
}

func generate(p plot) {

	// notify wait group
	wg.Add(1)
	defer wg.Done()

	// copy and modify the jobs according to the configuration
	var copiedJobs []job

	// iterate over all data collection jobs
	for _, j := range p.jobs {

		var i = j.ident
		if j.ident == "" {
			i = p.lineIdent
		}

		// determine name for the variable
		var n = j.name
		if p.forceName || j.name == "" {
			n = strings.Join(j.fileKeywords, "_")
		}

		// make a deep copy of the jobs
		// because the name modification by the force flag can happen in parallel
		copiedJobs = append(copiedJobs, job{
			name:                n,
			fileKeywords:        j.fileKeywords,
			excludeFileKeywords: j.excludeFileKeywords,
			ident:               i,
		})
	}

	// init data store for this plot
	var data = make(map[string][]float64)

	// open input directory
	files, err := ioutil.ReadDir("experiment_logs")
	if err != nil {
		log.Fatal(err)
	}

	// apply jobs to all files
	for _, f := range files {
		parseJobs(p, data, f, p.reg, p.lineIdent, copiedJobs)
	}

	// generate output python script
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

	// parse python template
	t, err := template.ParseFiles(p.script)
	if err != nil {
		log.Fatal(err)
	}

	// create output script path
	outScriptPath := filepath.Dir(p.script) + "/generated/" + strings.TrimSuffix(filepath.Base(p.script), ".py") + ".py"
	f, err := os.Create(outScriptPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// generate python log statements that log the data to stdout
	var logStatements = "print(\"script: " + p.script + "\")\n"
	for k := range data {
		logStatements += "print(\"stats.mean(" + k + "):\", stats.mean(" + k + "))\n"
	}

	// generate list of objects and loading instructions
	var (
		current  int
		load     string
		objects  string
		lastElem = len(data)
	)
	for _, j := range copiedJobs {
		if _, ok := data[j.name]; ok {

			current++
			if current == lastElem {
				objects += "'" + strings.ReplaceAll(j.name, "_", " ") + "'"
				if p.mean {
					load += "\tstats.mean(" + j.name + ")"
				} else {
					load += "\t" + j.name + ""
				}
			} else {
				objects += "'" + strings.ReplaceAll(j.name, "_", " ") + "',"
				if p.mean {
					load += "\tstats.mean(" + j.name + "),\n"
				} else {
					load += "\t" + j.name + ",\n"
				}
			}
		} else {
			fmt.Println(p.script+": no data for", j.name, "found after collecting values")
		}
	}

	// in stacked mode, update the loading instructions
	if p.stacked {
		// TODO: generate this properly and make it configurable
		load = "web = [\n" + load + "\n]\n"
	}

	// determine name for output image
	outImageName := filepath.Dir(p.script) + "/images/" + strings.TrimSuffix(filepath.Base(p.script), ".py") + ".png"

	// execute template with the generated data
	err = t.Execute(f, struct {
		Data    string
		Log     string
		Load    string
		Objects string
		Out     string
	}{
		out,
		logStatements,
		load,
		objects,
		"'" + outImageName + "'",
	})
	if err != nil {
		log.Fatal(err)
	}

	// execute the python script to generate the plot image
	outCmd, err := exec.Command("python3", outScriptPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(outCmd))
		log.Fatal(err)
	}

	if *flagDebug {
		fmt.Println(string(outCmd))
	}

	fmt.Println("created", outImageName)
}

func parseJobs(p plot, data map[string][]float64, f os.FileInfo, reg *regexp.Regexp, lineIdent string, jobs []job) {

	if reg == nil {
		reg = defaultRegEx
	}

	for _, j := range jobs {

		if j.ident == "" {
			j.ident = lineIdent
		}

		var nextJob bool

		// make sure all file keywords are part of the file name
		for _, w := range j.fileKeywords {
			if !strings.Contains(f.Name(), w) {
				nextJob = true
				break
			}
		}

		// make sure none of the excluded ones is
		for _, w := range j.excludeFileKeywords {
			if strings.Contains(f.Name(), w) {
				nextJob = true
				break
			}
		}

		if nextJob {
			continue
		}

		// TODO: read every file only once and cache in memory
		contents, err := ioutil.ReadFile("experiment_logs/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		// iterate over file contents line by line
		for _, l := range strings.Split(string(contents), "\n") {

			if !strings.Contains(l, j.ident) {
				continue
			}

			// apply regular expression to grab the data of interest
			value := reg.FindString(l)

			// remove the XXX= prefix from the extracted value
			for i, c := range value {
				if c == '=' {
					value = value[i+1:]
				}
			}

			if *flagDebug {
				fmt.Println(ansi.Red+filepath.Base(p.script)+": "+ansi.Yellow+j.name+": "+ansi.Blue+f.Name()+":"+ansi.Reset, value, l)
			}

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

			// parse value as float
			f, err := strconv.ParseFloat(
				final,
				64)
			if err != nil {
				log.Fatal("failed to parse float ", final, ", error: ", err)
			}

			// convert seconds to millis if necessary
			if toMs {
				f = f * 1000
			}

			// add to data store
			data[j.name] = append(data[j.name], f)

			// next iteration
			continue
		}
	}
}
