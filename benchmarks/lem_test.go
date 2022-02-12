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

package benchmarks_test

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
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
	cout  *string
	alloc int64
	bytes int64
	regxp []*regexp.Regexp
}

func (tc lemTestCase) run(t *testing.T) {
	// Assert the expected leak, escape, move decisions match.
	em := make(map[int]struct{}, len(tc.regxp))
	for k, rx := range tc.regxp {
		if rx.FindString(*tc.cout) != "" {
			em[k] = struct{}{}
		}
	}
	for k, rx := range tc.regxp {
		if _, ok := em[k]; !ok {
			t.Errorf("exp.m=%s", rx)
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

func compile(file string) (string, error) {
	b, err := exec.Command("go", "tool", "compile", "-m", file).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getTestCaseGroups() ([]lemTestCaseGroup, error) {

	// Get the path to this source file.
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("cannot get filename")
	}

	// Compile this source file & print optimizations output.
	compileOutput, err := compile(fileName)
	if err != nil {
		return nil, err
	}

	var (
		lineNo     = 1
		lemGroupX  int
		lemTCaseX  int
		lemGroups  = []lemTestCaseGroup{}
		lemNameRx  = regexp.MustCompile(`^// lem\.((\w+)\d+)\.name=(.+)$`)
		lemAllocRx = regexp.MustCompile(`^// lem\.[^.]+\.alloc=(.+)$`)
		lemBytesRx = regexp.MustCompile(`^// lem\.[^.]+\.bytes=(.+)$`)
		lemMatchRx = regexp.MustCompile(`// lem\.[^.]+\.m=(.+)$`)
	)

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if m := lemNameRx.FindStringSubmatch(line); m != nil {
			if len(lemGroups) == 0 || lemGroups[lemGroupX].name != m[2] {
				lemGroups = append(lemGroups, lemTestCaseGroup{
					name: m[2],
				})
				lemGroupX = len(lemGroups) - 1
				lemTCaseX = 0
			}
			lemGroups[lemGroupX].testCases = append(
				lemGroups[lemGroupX].testCases,
				lemTestCase{
					name: m[3],
					fvnc: lemFuncs[m[1]],
					cout: &compileOutput,
				},
			)
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
			r, err := regexp.Compile(fmt.Sprintf("%d:\\d+: %s", lineNo, m[1]))
			if err != nil {
				return nil, err
			}
			lemGroups[lemGroupX].testCases[lemTCaseX].regxp = append(
				lemGroups[lemGroupX].testCases[lemTCaseX].regxp, r)
		}
		lineNo++
	}

	return lemGroups, nil
}

func eqBytes(exp, act int64) bool {
	if exp%oneMiB == 0 {
		return (exp / oneMiB) == (act / oneMiB)
	}
	return exp == act
}

// lem.leak1.name=to sink
// lem.leak1.alloc=0
// lem.leak1.bytes=0
func leak1(b *testing.B) {
	var sink *int32
	f := func(p *int32) { // lem.leak1.m=leaking param: p
		sink = p
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(nil)
	}
	_ = sink
}
func init() {
	lemFuncs["leak1"] = leak1
}

// lem.leak2.name=to result
// lem.leak2.alloc=0
// lem.leak2.bytes=0
func leak2(b *testing.B) {
	f := func(p *int32) *int32 { // lem.leak2.m=leaking param: p to result ~r[0-1] level=0
		return p
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(nil)
	}
}
func init() {
	lemFuncs["leak2"] = leak2
}

// lem.noescape1.name=pointer does not outlive its call stack
// lem.noescape1.alloc=0
// lem.noescape1.bytes=0
func noescape1(b *testing.B) {
	f := func(p *int32) *int32 { // lem.noescape1.m=leaking param: p to result ~r[0-1] level=0
		return p
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(new(int32)) // lem.noescape1.m=new\(int32\) does not escape
	}
}
func init() {
	lemFuncs["noescape1"] = noescape1
}

// lem.escape1.name=pointer outlived its call stack
// lem.escape1.alloc=1
// lem.escape1.bytes=4
func escape1(b *testing.B) {
	var sink *int32
	f1 := func() *int32 {
		return new(int32)
	}
	f2 := func(p *int32) *int32 { // lem.escape1.m=leaking param: p to result ~r[0-1] level=0
		return p
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = (f2(f1())) // lem.escape1.m=new\(int32\) escapes to heap
	}
	_ = sink
}
func init() {
	lemFuncs["escape1"] = escape1
}

// lem.escape2.name=heap cannot point to stack
// lem.escape2.alloc=1
// lem.escape2.bytes=4
func escape2(b *testing.B) {
	var sink *int32
	f := func(p *int32) *int32 { // lem.escape1.m=leaking param: p to result ~r[0-1] level=0
		return p
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = f(new(int32)) // lem.escape1.m=new\(int32\) escapes to heap
	}
	_ = sink
}
func init() {
	lemFuncs["escape2"] = escape2
}

// lem.escape3.name=heap cannot point to stack; no malloc bc iface & zero value
// lem.escape3.alloc=0
// lem.escape3.bytes=0
func escape3(b *testing.B) {
	var sink interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32
		sink = x // lem.escape3.m=x escapes to heap
	}
	_ = sink
}
func init() {
	lemFuncs["escape3"] = escape3
}

// lem.escape4.name=heap cannot point to stack; no malloc bc iface & byte
// lem.escape4.alloc=0
// lem.escape4.bytes=0
func escape4(b *testing.B) {
	var sink interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x byte
		x = 253
		sink = x // lem.escape4.m=x escapes to heap
	}
	_ = sink
}
func init() {
	lemFuncs["escape4"] = escape4
}

// lem.escape5.name=heap cannot point to stack; no malloc bc iface & single byte-wide value
// lem.escape5.alloc=0
// lem.escape5.bytes=0
func escape5(b *testing.B) {
	var sink interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int64
		x = 253
		sink = x // lem.escape5.m=x escapes to heap
	}
	_ = sink
}
func init() {
	lemFuncs["escape5"] = escape5
}

// lem.move1.name=pointing to stack from heap
// lem.move1.alloc=1
// lem.move1.bytes=4
func move1(b *testing.B) {
	var sink *int32
	for i := 0; i < b.N; i++ {
		var x int32 = 4096 // lem.move1.m=moved to heap: x
		sink = &x
	}
	_ = sink
}
func init() {
	lemFuncs["move1"] = move1
}

// lem.move2.name=too large for stack
// lem.move2.alloc=1
// lem.move2.bytes=15728640
func move2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf [15 * 1024 * 1024]byte // lem.move2.m=moved to heap: buf
		_ = buf
	}
}
func init() {
	lemFuncs["move2"] = move2
}
