/*
Copyright 2022

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lem_test

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"testing"
)

func TestLem(t *testing.T) {
	if f := flag.Lookup("test.benchtime"); f != nil {
		og := f.Value.String()
		f.Value.Set("100x")
		defer func() {
			f.Value.Set(og)
		}()
	}
	if f := flag.Lookup("test.benchmem"); f != nil {
		og := f.Value.String()
		f.Value.Set("true")
		defer func() {
			f.Value.Set(og)
		}()
	}

	// Get the test case groups.
	testCaseGroups, err := getTestCaseGroups()
	if err != nil {
		t.Fatal(err)
	}

	// Run the test case groups.
	for i := range testCaseGroups {
		t.Run(testCaseGroups[i].name, testCaseGroups[i].run)
	}
}

var (
	lemFuncs = map[string]func(*testing.B){}
)

const (
	oneMiB     = 1024 * 1024
	fifteenMiB = 15 * oneMiB
)

type lemMatch struct {
	lineNo int
}

type lemTestCaseGroup struct {
	name      string
	testCases []lemTestCase
}

func (tcg lemTestCaseGroup) run(t *testing.T) {
	for i := range tcg.testCases {
		t.Run(tcg.testCases[i].name, tcg.testCases[i].run)
	}
}

type lemTestCase struct {
	name  string
	fvnc  func(*testing.B)
	bout  *string
	alloc int64
	bytes int64
	match []*regexp.Regexp
	natch []*regexp.Regexp
}

func (tc lemTestCase) run(t *testing.T) {
	// Assert the expected leak, escape, move decisions match.
	for _, rx := range tc.match {
		if rx.FindString(*tc.bout) == "" {
			t.Errorf("exp.m=%s", rx)
		}
	}

	// Assert the expected leak, escape, move decisions do not match.
	for _, rx := range tc.natch {
		if rx.FindString(*tc.bout) != "" {
			t.Errorf("exp.m!=%s", rx)
		}
	}

	// Assert the expected allocs and bytes match.
	r := testing.Benchmark(tc.fvnc)
	if ea, aa := tc.alloc, r.AllocsPerOp(); ea != aa {
		t.Errorf("exp.alloc=%d, act.alloc=%d", ea, aa)
	}
	if eb, ab := tc.bytes, r.AllocedBytesPerOp(); !eqBytes(eb, ab) {
		t.Errorf("exp.bytes=%d, act.bytes=%d", eb, ab)
	}
}

func build(dir string) (string, error) {
	var s string
	b, err := exec.Command(
		"go", "test", "-c", "-gcflags", "-m", dir).CombinedOutput()
	if len(b) > 0 {
		s = string(b)
	}
	if err != nil {
		fmt.Println(s)
		return "", err
	}
	return s, nil
}

func getTestCaseGroups() ([]lemTestCaseGroup, error) {
	// Get the path to this source file.
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("cannot get filename")
	}

	// Get the path to the directory of the source to know
	// what package to build.
	dirPath := filepath.Dir(filePath)

	// Build the package and print the optimization output.
	buildOutput, err := build(dirPath)
	if err != nil {
		return nil, err
	}

	var (
		lemGroupX  int
		lemTCaseX  int
		lemGroups  = []lemTestCaseGroup{}
		lemNameRx  = regexp.MustCompile(`^// lem\.((\w+?)\d+)\.name=(.+)$`)
		lemAllocRx = regexp.MustCompile(`^// lem\.[^.]+\.alloc=(.+)$`)
		lemBytesRx = regexp.MustCompile(`^// lem\.[^.]+\.bytes=(.+)$`)
		lemMatchRx = regexp.MustCompile(`// lem\.[^.]+\.m=(.+)$`)
		lemNatchRx = regexp.MustCompile(`// lem\.[^.]+\.m!=(.+)$`)
	)

	// Get a list of the test sources.
	srcFilePaths, err := filepath.Glob(path.Join(dirPath, "lem_*_test.go"))
	if err != nil {
		return nil, err
	}

	for _, filePath := range srcFilePaths {
		var (
			lines    []string
			lineNo   = 1
			fileName = filepath.Base(filePath)
		)

		f, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
			if m := lemNameRx.FindStringSubmatch(line); m != nil {
				if len(lemGroups) == 0 || lemGroups[lemGroupX].name != m[2] {
					lemGroups = append(lemGroups, lemTestCaseGroup{
						name: m[2],
					})
					lemGroupX = len(lemGroups) - 1
					lemTCaseX = 0
				}
				ltc := lemTestCase{
					name: m[3],
					fvnc: lemFuncs[m[1]],
					bout: &buildOutput,
				}
				lemGroups[lemGroupX].testCases = append(
					lemGroups[lemGroupX].testCases, ltc)
				lemTCaseX = len(lemGroups[lemGroupX].testCases) - 1
			} else if m := lemAllocRx.FindStringSubmatch(line); m != nil {
				i, err := strconv.ParseInt(m[1], 10, 64)
				if err != nil {
					return nil, err
				}
				lemGroups[lemGroupX].testCases[lemTCaseX].alloc = i
			} else if m := lemBytesRx.FindStringSubmatch(line); m != nil {
				i, err := strconv.ParseInt(m[1], 10, 64)
				if err != nil {
					return nil, err
				}
				lemGroups[lemGroupX].testCases[lemTCaseX].bytes = i
			} else if m := lemMatchRx.FindStringSubmatch(line); m != nil {
				r, err := regexp.Compile(
					fmt.Sprintf("%s:%d:\\d+: %s", fileName, lineNo, m[1]),
				)
				if err != nil {
					return nil, err
				}
				lemGroups[lemGroupX].testCases[lemTCaseX].match = append(
					lemGroups[lemGroupX].testCases[lemTCaseX].match, r)
			} else if m := lemNatchRx.FindStringSubmatch(line); m != nil {
				r, err := regexp.Compile(
					fmt.Sprintf("%s:%d:\\d+: %s", fileName, lineNo, m[1]),
				)
				if err != nil {
					return nil, err
				}
				lemGroups[lemGroupX].testCases[lemTCaseX].natch = append(
					lemGroups[lemGroupX].testCases[lemTCaseX].natch, r)
			}
			lineNo++
		}
	}
	return lemGroups, nil
}

func eqBytes(exp, act int64) bool {
	if exp != 0 && exp%1024 == 0 {
		return (exp / oneMiB) == (act / oneMiB)
	}
	return exp == act
}
