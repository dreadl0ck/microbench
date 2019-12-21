package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"text/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile("delta=[0-9]*.[0-9]*m?s")

type job struct {
	fileKeywords        []string
	excludeFileKeywords []string
	lineIdent           string
	name                string
}

func main() {
	// mean-hashing-time-sequential.png
	generate(
		"plots/scripts/mean-hashing-time-sequential.py",
		"hash loop benchmark",
		&job{
			fileKeywords:        []string{"qemu", "sequential"},
			excludeFileKeywords: []string{"emulated"},
			name: "qemu",
		},
		&job{
			fileKeywords:        []string{"qemu", "sequential", "emulated"},
			name: "qemu_emulated",
		},
		&job{
			fileKeywords:        []string{"firecracker", "sequential"},
			name: "firecracker",
		},
	)

	// mean-hashing-time-concurrent.png
	generate(
		"plots/scripts/mean-hashing-time-concurrent.py",
		"hash loop benchmark",
		&job{
			fileKeywords:        []string{"qemu", "x10"},
			excludeFileKeywords: []string{"emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x20"},
			excludeFileKeywords: []string{"emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x10", "emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x20", "emulated"},
		},
		&job{
			fileKeywords:        []string{"firecracker", "x10"},
		},
		&job{
			fileKeywords:        []string{"firecracker", "x20"},
		},
	)

	// mean-kernel-boot-time-sequential.png
	generate(
		"plots/scripts/mean-kernel-boot-time-sequential.py",
		"kernel boot time received",
		&job{
			fileKeywords:        []string{"qemu", "sequential"},
			excludeFileKeywords: []string{"emulated"},
			name: "qemu",
		},
		&job{
			fileKeywords:        []string{"qemu", "sequential", "emulated"},
			name: "qemu_emulated",
		},
		&job{
			fileKeywords:        []string{"firecracker", "sequential"},
			name: "firecracker",
		},
	)

	// mean-kernel-boot-time-concurrent.png
	generate(
		"plots/scripts/mean-kernel-boot-time-concurrent.py",
		"kernel boot time received",
		&job{
			fileKeywords:        []string{"qemu", "x10"},
			excludeFileKeywords: []string{"emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x20"},
			excludeFileKeywords: []string{"emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x10", "emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x20", "emulated"},
		},
		&job{
			fileKeywords:        []string{"firecracker", "x10"},
		},
		&job{
			fileKeywords:        []string{"firecracker", "x20"},
		},
	)

	// mean-webservice-time-sequential.png
	generate(
		"plots/scripts/mean-webservice-time-sequential.py",
		"time until HTTP reply from webservice",
		&job{
			fileKeywords:        []string{"qemu", "sequential"},
			excludeFileKeywords: []string{"emulated"},
			name: "qemu",
		},
		&job{
			fileKeywords:        []string{"qemu", "sequential", "emulated"},
			name: "qemu_emulated",
		},
		&job{
			fileKeywords:        []string{"firecracker", "sequential"},
			name: "firecracker",
		},
	)

	// mean-webservice-time-concurrent.png
	generate(
		"plots/scripts/mean-webservice-time-concurrent.py",
		"time until HTTP reply from webservice",
		&job{
			fileKeywords:        []string{"qemu", "x10"},
			excludeFileKeywords: []string{"emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x20"},
			excludeFileKeywords: []string{"emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x10", "emulated"},
		},
		&job{
			fileKeywords:        []string{"qemu", "x20", "emulated"},
		},
		&job{
			fileKeywords:        []string{"firecracker", "x10"},
		},
		&job{
			fileKeywords:        []string{"firecracker", "x20"},
		},
	)
}

func generate(script string, lineIdent string, jobs ...*job) {

	for _, j := range jobs {
		if j.name == "" {
			j.name = strings.Join(j.fileKeywords, "_")
		}
	}

	var data = make(map[string][]float64)

	files, err := ioutil.ReadDir("experiment_logs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		parseJobs(data, f, lineIdent, jobs)
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

	fmt.Println(out)

	t, err := template.ParseFiles(script)
	if err != nil {
		log.Fatal(err)
	}

	outScriptPath := filepath.Dir(script) + "/generated/" + strings.TrimSuffix(filepath.Base(script), ".py") + ".py"
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
	for _, j := range jobs {
		if _, ok := data[j.name]; ok {

			current++
			if current == lastElem {
				objects += "'" + strings.ReplaceAll(j.name, "_", " ") + "'"
				load += "\tstats.mean("+ j.name + ")"
			} else {
				objects += "'" + strings.ReplaceAll(j.name, "_", " ") + "',"
				load += "\tstats.mean("+ j.name + "),\n"
			}
		} else {
			log.Fatal("name not found: ", j.name)
		}
	}

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
		"'"+ filepath.Dir(script) + "/images/" + strings.TrimSuffix(filepath.Base(script), ".py") + ".png'",
	})
	if err != nil {
		log.Fatal(err)
	}

	outCmd, err := exec.Command("python3", outScriptPath).CombinedOutput()
	if err != nil {
		fmt.Println(string(outCmd))
		log.Fatal(err)
	}
	fmt.Println(string(outCmd))
}

func parseJobs(data map[string][]float64, f os.FileInfo, lineIdent string, jobs []*job) {

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

			value := lineRegex.FindString(l)
			value = strings.TrimPrefix(value, "delta=")

			fmt.Println(f.Name(), value)

			var toMs bool
			if !strings.HasSuffix(value, "ms") {
				toMs = true
			}
			final := strings.TrimSuffix(
				strings.TrimSuffix(value, "ms"),
				"s")

			fmt.Println(final)

			f, err := strconv.ParseFloat(
				final,
				64)
			if err != nil {
				log.Fatal(err)
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
