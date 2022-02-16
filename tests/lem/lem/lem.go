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

package lem

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// MyDirectory returns the filesystem path to the directory of the caller.
//
// This function is useful for callers wanting to populate a B.IncldueDirs
// field.
func MyDirectory() string {
	_, callersFilePath, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Dir(callersFilePath)
}

// Sets the value of the -test.benchtime flag and returns the original
// value if one was present, otherwise an empty string is returned.
//
// Please note this function is a no-op if the flag is not already
// defined.
func SetBenchtime(s string) string {
	f := flag.Lookup("test.benchtime")
	if f == nil {
		return ""
	}
	og := f.Value.String()
	f.Value.Set(s)
	return og
}

// Sets the value of the -test.benchmem flag and returns the original
// value if one was present, otherwise an empty string is returned.
//
// Please note this function is a no-op if the flag is not already
// defined.
func SetBenchmem(s string) string {
	f := flag.Lookup("test.benchmem")
	if f == nil {
		return ""
	}
	og := f.Value.String()
	f.Value.Set(s)
	return og
}

// B is a test harness for asserting when values leak, escape, or move to
// the heap and the number of allocations and memory as a result.
type B struct {
	// Benchmarks is a map of named, benchmark functions.
	//
	// The keys in this map should correspond to the <ID> from
	// "lem.<ID>" comments.
	Benchmarks map[string]func(*testing.B)

	// IncludeDirs is a list of directories that are scanned for
	// test sources and built using the compiler optimization flag.
	IncludeDirs []string
}

// Run executes the tests.
func (bb B) Run(t *testing.T) {
	// Build all the included directories.
	buildOutput, err := GoBuildTest(bb.IncludeDirs...)
	if err != nil {
		t.Fatal(err)
	}

	testCases, err := GetTestCases(bb.IncludeDirs...)
	if err != nil {
		t.Fatal(err)
	}

	NewTestCaseTree(testCases...).Run(t, buildOutput, bb.Benchmarks)
}

// TestCaseTree organizes the TestCases in a tree structure.
type TestCaseTree struct {
	sync.Once
	Index map[string]int
	Steps []string
	Nodes []TestCaseTree
	Tests []TestCase
}

// Insert a TestCase into the tree.
func (tr *TestCaseTree) Insert(testCase TestCase, path ...string) {
	tr.Once.Do(func() { tr.Index = map[string]int{} })
	if len(path) < 2 {
		if len(path) == 1 {
			testCase.Name = path[0]
		}
		tr.Tests = append(tr.Tests, testCase)
	} else {
		if _, ok := tr.Index[path[0]]; !ok {
			tr.Index[path[0]] = len(tr.Nodes)
			tr.Nodes = append(tr.Nodes, TestCaseTree{})
			tr.Steps = append(tr.Steps, path[0])
		}
		tr.Nodes[tr.Index[path[0]]].Insert(testCase, path[1:]...)
	}
}

// Run the tests for this tree.
func (tr TestCaseTree) Run(
	t *testing.T,
	buildOutput string,
	benchmarks map[string]func(*testing.B)) {

	// Descend into any possible children.
	for i, s := range tr.Steps {
		i, s := i, s
		t.Run(s, func(t *testing.T) {
			tr.Nodes[i].Run(t, buildOutput, benchmarks)
		})
	}

	// Run this node's tests.
	for i := range tr.Tests {
		tc := tr.Tests[i]

		t.Run(tc.Name, func(t *testing.T) {
			// Assert the expected leak, escape, move decisions match.
			for _, rx := range tc.Matches {
				if rx.FindString(buildOutput) == "" {
					t.Errorf("exp.m=%s", rx)
				}
			}

			// Assert the expected leak, escape, move decisions do not match.
			for _, rx := range tc.Natches {
				if s := rx.FindString(buildOutput); s != "" {
					t.Errorf("exp.m!=%s, found=%s", rx, s)
				}
			}

			// Find the benchmark function.
			if benchFn, ok := benchmarks[tc.ID]; !ok {
				t.Errorf("benchmark function not registered for %s", tc.ID)
			} else {
				// Assert the expected allocs and bytes match.
				r := testing.Benchmark(benchFn)
				if ea, aa := tc.AllocOp, r.AllocsPerOp(); !ea.Eq(aa) {
					t.Errorf("exp.alloc=%d, act.alloc=%d", ea, aa)
				}
				if eb, ab := tc.BytesOp, r.AllocedBytesPerOp(); !eb.Eq(ab) {
					t.Errorf("exp.bytes=%d, act.bytes=%d", eb, ab)
				}
			}
		})
	}
}

// NewTestCaseTree returns a new TestCaseTree with the provided test cases.
func NewTestCaseTree(testCases ...TestCase) TestCaseTree {
	var tree TestCaseTree
	for _, tc := range testCases {
		tree.Insert(tc, GetTestCasePath(tc.ID, tc.Name)...)
	}
	return tree
}

// TestCase is a test case parsed from the lem comments in a source file.
type TestCase struct {
	// ID maps to lem.<ID>.
	ID string

	// Name maps to lem.<ID>.name=<NAME>.
	// Please see the package documentation for more information.
	Name string

	// AllocOp is the expected number of allocations per operation.
	AllocOp Int64Range

	// BytesOp is the expected number of bytes per op.
	BytesOp Int64Range

	// Matches is a list of patterns that must appear in the optimization
	// output.
	Matches []*regexp.Regexp

	// Natches is a list of patterns that must not appear in the optimization
	// output.
	Natches []*regexp.Regexp
}

// Int64Range is an inclusive range of int64 values.
type Int64Range struct {
	Min int64
	Max int64
}

// Eq returns true if a when (Min==Max && a==Min) || (a>=Min && a<=Max).
func (i Int64Range) Eq(a int64) bool {
	if i.Min == i.Max {
		return i.Min == a
	}
	return a >= i.Min && a <= i.Max
}

// String returns the string version of this value.
func (i Int64Range) String() string {
	if i.Min == i.Max {
		return fmt.Sprintf("%d", i.Min)
	}
	return fmt.Sprintf("%d-%d", i.Min, i.Max)
}

// GoBuildTest uses "go test -c -gcflags -m" to build the test packages at the
// specified directories and return the optimizations.
func GoBuildTest(dirs ...string) (string, error) {
	var stderr bytes.Buffer
	for _, d := range dirs {
		var stderr2 bytes.Buffer
		stderr3 := io.MultiWriter(&stderr, &stderr2)
		cmd := exec.Command(
			"go", "test", "-v", "-count", "1", "-c", "-gcflags", "-m", d)
		cmd.Stderr = stderr3
		if err := cmd.Run(); err != nil {
			return "", fmt.Errorf("%w\n%s", err, stderr2.String())
		}
	}
	return stderr.String(), nil
}

// GetTestCasePath returns the test case path name from the provided
// ID and name.
// Please see the package documentation for more information.
func GetTestCasePath(id, name string) []string {
	var path []string

	if len(name) == 0 || name[0] != '/' {
		// If the ID ends with an integer then treat the part before the
		// integer and the integer itself as separate path elements.
		if m := intSuffixRx.FindStringSubmatch(id); m != nil {
			path = append(path, m[1], m[2])
		} else {
			path = append(path, id)
		}
	}
	if len(name) > 0 {
		path = append(path, strings.Split(name, "/")...)
	}

	// Remove any empty elements from the slice.
	temp := path[:0]
	for _, s := range path {
		if s != "" {
			temp = append(temp, s)
		}
	}

	return temp
}

var (
	nameRx      = regexp.MustCompile(`(?m)^// lem\.([^.]+)\.name=(.+)$`)
	allocRx     = regexp.MustCompile(`(?m)^// lem\.([^.]+)\.alloc=(\d+)(?:-(\d+))?$`)
	bytesRx     = regexp.MustCompile(`(?m)^// lem\.([^.]+)\.bytes=(\d+)(?:-(\d+))?$`)
	matchRx     = regexp.MustCompile(`(?m)// lem\.([^.]+)\.m=(.+)$`)
	natchRx     = regexp.MustCompile(`(?m)// lem\.([^.]+)\.m!=(.+)$`)
	intSuffixRx = regexp.MustCompile(`^(\w+?)(\d+)$`)
)

// TestCaseLookupTable provides a quick way to check if a test case already
// exists.
type TestCaseLookupTable map[string]*TestCase

// Get the test case with the specified ID, otherwise an error is returned.
func (t TestCaseLookupTable) Get(id string) (*TestCase, error) {
	tc, ok := t[id]
	if !ok {
		return nil, fmt.Errorf("unknown test case ID: %s", id)
	}
	return tc, nil
}

// GetTestCases returns a list of test cases scanned from test sources in the
// provided directories.
func GetTestCases(dirs ...string) ([]TestCase, error) {

	var (
		testCases []TestCase
		lookupTbl = TestCaseLookupTable{}
	)

	for _, d := range dirs {
		// Get all of the test files in the provided directory.
		files, err := filepath.Glob(path.Join(d, "*_test.go"))
		if err != nil {
			return nil, err
		}
		for _, filePath := range files {
			testCasesInFile, err := GetTestCasesInFile(filePath, lookupTbl)
			if err != nil {
				return nil, err
			}

			// Store the length of the testCases slice and then append the
			// test cases from the file to it.
			indexOfUnmappedTestCases := len(testCases)
			testCases = append(testCases, testCasesInFile...)

			// Add the newly appended test cases to the lookup table.
			for i := indexOfUnmappedTestCases; i < len(testCases); i++ {
				lookupTbl[testCases[i].ID] = &testCases[i]
			}
		}
	}
	return testCases, nil
}

func GetTestCasesInFile(
	filePath string,
	lookupTbl TestCaseLookupTable) ([]TestCase, error) {

	var (
		testCases []TestCase
		lineNo    = 1
		fileName  = filepath.Base(filePath)
	)

	if lookupTbl == nil {
		lookupTbl = TestCaseLookupTable{}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Scan each line of the file for lem comments.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()

		// lem.<ID>.name=<NAME>
		if m := nameRx.FindStringSubmatch(l); m != nil {
			id := m[1]
			if tc, _ := lookupTbl.Get(id); tc != nil {
				return nil, fmt.Errorf("duplicate test case ID: %s", id)
			}
			testCases = append(testCases, TestCase{ID: id, Name: m[2]})
			lookupTbl[id] = &testCases[len(testCases)-1]
		} else if m := allocRx.FindStringSubmatch(l); m != nil {
			tc, err := lookupTbl.Get(m[1])
			if err != nil {
				return nil, err
			}
			min, err := strconv.ParseInt(m[2], 10, 64)
			if err != nil {
				return nil, err
			}
			tc.AllocOp.Min = min
			if len(m) < 3 || m[3] == "" {
				tc.AllocOp.Max = min
			} else {
				max, err := strconv.ParseInt(m[3], 10, 64)
				if err != nil {
					return nil, err
				}
				tc.AllocOp.Max = max
			}
		} else if m := bytesRx.FindStringSubmatch(l); m != nil {
			tc, err := lookupTbl.Get(m[1])
			if err != nil {
				return nil, err
			}
			min, err := strconv.ParseInt(m[2], 10, 64)
			if err != nil {
				return nil, err
			}
			tc.BytesOp.Min = min
			if len(m) < 3 || m[3] == "" {
				tc.BytesOp.Max = min
			} else {
				max, err := strconv.ParseInt(m[3], 10, 64)
				if err != nil {
					return nil, err
				}
				tc.BytesOp.Max = max
			}
		} else if m := matchRx.FindStringSubmatch(l); m != nil {
			tc, err := lookupTbl.Get(m[1])
			if err != nil {
				return nil, err
			}
			r, err := regexp.Compile(
				fmt.Sprintf("%s:%d:\\d+: %s", fileName, lineNo, m[2]),
			)
			if err != nil {
				return nil, err
			}
			tc.Matches = append(tc.Matches, r)
		} else if m := natchRx.FindStringSubmatch(l); m != nil {
			tc, err := lookupTbl.Get(m[1])
			if err != nil {
				return nil, err
			}
			r, err := regexp.Compile(
				fmt.Sprintf("%s:%d:\\d+: %s", fileName, lineNo, m[2]),
			)
			if err != nil {
				return nil, err
			}
			tc.Natches = append(tc.Natches, r)
		}

		lineNo++
	}

	return testCases, nil
}
