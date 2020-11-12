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
	"github.com/dreadl0ck/microbench/stats"
	"regexp"
	"strings"
	"time"
)

var tsRegExp = regexp.MustCompile("[0-9]*\\.[0-9]*")

func parseKernelLog(contents []byte) (s *stats.Summary, err error) {

	s = &stats.Summary{}

	lines := strings.Split(string(contents), "\n")
	for i, line := range lines {
		if strings.Contains(line, "random: fast init done") {

			// ex: [   17.567752] random: fast init done
			duration := strings.TrimSpace(tsRegExp.FindString(lines[i])) + "s"
			l.Info("found duration:", duration)

			d, err := time.ParseDuration(duration)
			if err != nil {
				return nil, err
			}

			s.KernelBootup = d
			s.KernelLogLines = i + 1
			//s.KernelLogs = string(contents)
			break
		}
	}
	return
}
