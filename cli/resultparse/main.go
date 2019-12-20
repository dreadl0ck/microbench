package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var data = make(map[string][]float64)

func main() {

	files, err := ioutil.ReadDir("experiment_logs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		parse(f, []string{"qemu", "sequential"}, []string{"emulated"}, []string{"hash loop benchmark"}, regexp.MustCompile("[0-9]*.[0-9]*m?s,"), "qemu_sequential")
		parse(f, []string{"qemu", "sequential-emulated-cpu"}, []string{}, []string{"hash loop benchmark"}, regexp.MustCompile("[0-9]*.[0-9]*m?s,"), "qemu_sequential_emulated")

		parse(f, []string{"qemu", "x10"}, []string{"emulated"}, []string{"hash loop benchmark"}, regexp.MustCompile("[0-9]*.[0-9]*m?s,"), "qemu_x10")
		parse(f, []string{"qemu", "x20"}, []string{"emulated"}, []string{"hash loop benchmark"}, regexp.MustCompile("[0-9]*.[0-9]*m?s,"), "qemu_x20")
		parse(f, []string{"qemu", "x10", "emulated"}, []string{}, []string{"hash loop benchmark"}, regexp.MustCompile("[0-9]*.[0-9]*m?s,"), "qemu_x10_emulated")
		parse(f, []string{"qemu", "x20", "emulated"}, []string{}, []string{"hash loop benchmark"}, regexp.MustCompile("[0-9]*.[0-9]*m?s,"), "qemu_x20_emulated")
	}

	var out string

	for key, values := range data {
		out += key + " = [\n"
		lastElem := len(values)-1
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
}

func parse(f os.FileInfo, fileKeywords []string, excludeFileKeywords []string, lineKeywords []string, lineRegex *regexp.Regexp, varname string)  {

	for _, w := range fileKeywords {
		if !strings.Contains(f.Name(), w) {
			return
		}
	}

	for _, w := range excludeFileKeywords {
		if strings.Contains(f.Name(), w) {
			return
		}
	}

	contents, err := ioutil.ReadFile("experiment_logs/"+f.Name())
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range strings.Split(string(contents), "\n") {
		var containsAll bool
		for _, lk := range lineKeywords {
			if strings.Contains(l, lk) {
				containsAll = true
			} else {
				containsAll = false
			}
		}
		if containsAll {

			value := lineRegex.FindString(l)
			fmt.Println(f.Name(), value)

			var toMs bool
			if !strings.HasSuffix(value, "ms,") {
				toMs = true
			}

			f, err := strconv.ParseFloat(
				strings.TrimSuffix(
					strings.TrimSuffix(value, "ms,"),
				"s,"),
			64)
			if err != nil {
				log.Fatal(err)
			}

			if toMs {
				f = f * 1000
			}

			data[varname] = append(data[varname], f)

			// next iteration
			continue
		}
	}
}