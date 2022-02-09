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
	"flag"
	"testing"
)

const (
	oneMiB     = 1024 * 1024
	fifteenMiB = 15 * oneMiB
)

//go:noinline
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

	testCases := []testCaseForBenchmark{
		{
			name: "no escape",
			fvnc: benchNoEscape,
		},
		{
			name: "move to heap",
			child: []testCaseForBenchmark{
				{
					name:  "pkg pointer to stack var",
					fvnc:  benchPkgPointerToStackVar,
					alloc: 1,
					bytes: 4,
				},
				{
					name:  "too large for stack frame",
					fvnc:  benchTooLargeForStackFrame,
					alloc: 1,
					bytes: fifteenMiB,
				},
			},
		},
		{
			name: "leak",
			child: []testCaseForBenchmark{
				{
					name: "to pkg var",
					child: []testCaseForBenchmark{
						{
							name:  "escape",
							fvnc:  benchEscapeByLeakingToPkgVar,
							alloc: 1,
							bytes: 4,
						},
					},
				},
				{
					name: "to result",
					child: []testCaseForBenchmark{
						{
							name: "store",
							child: []testCaseForBenchmark{
								{
									name: "none",
									fvnc: benchLeak2ResultStoreNone,
								},
								{
									name: "in loop",
									fvnc: benchLeak2ResultStoreInLoop,
								},
								{
									name:  "outside of loop",
									fvnc:  benchLeak2ResultStoreOutLoop,
									alloc: 1,
									bytes: 4,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "chan",
			child: []testCaseForBenchmark{
				{
					name: "outside of loop",
					child: []testCaseForBenchmark{
						{
							name: "int32",
							child: []testCaseForBenchmark{
								{
									name: "val",
									child: []testCaseForBenchmark{
										{
											name: "nzd",
											fvnc: benchChanInt32,
										},
										{
											name: "zed",
											fvnc: benchChanInt32Zed,
										},
									},
								},
								{
									name: "ptr",
									child: []testCaseForBenchmark{
										{
											name:  "nzd",
											fvnc:  benchChanPtrInt32,
											alloc: 1,
											bytes: 4,
										},
										{
											name: "zed",
											fvnc: benchChanPtrInt32Zed,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "iface",
			child: []testCaseForBenchmark{
				{
					name: "in loop",
					child: []testCaseForBenchmark{
						{
							name: "int32",
							child: []testCaseForBenchmark{
								{
									name: "val",
									child: []testCaseForBenchmark{
										{
											name: "nzd",
											fvnc: benchIfaceInLoopInt32,
										},
										{
											name: "zed",
											fvnc: benchIfaceInLoopInt32Zed,
										},
									},
								},
								{
									name: "ptr",
									child: []testCaseForBenchmark{
										{
											name: "nzd",
											fvnc: benchIfaceInLoopPtrInt32,
										},
										{
											name: "zed",
											fvnc: benchIfaceInLoopPtrInt32Zed,
										},
									},
								},
							},
						},
					},
				},
				{
					name: "outside of loop",
					child: []testCaseForBenchmark{
						{
							name: "int32",
							child: []testCaseForBenchmark{
								{
									name: "val",
									child: []testCaseForBenchmark{
										{
											name:  "nzd",
											fvnc:  benchIfaceOutLoopInt32,
											alloc: 1,
											bytes: 4,
										},
										{
											name: "zed",
											fvnc: benchIfaceOutLoopInt32Zed,
										},
									},
								},
								{
									name: "ptr",
									child: []testCaseForBenchmark{
										{
											name:  "nzd",
											fvnc:  benchIfaceOutLoopPtrInt32,
											alloc: 1,
											bytes: 4,
										},
										{
											name: "zed",
											fvnc: benchIfaceOutLoopPtrInt32Zed,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "slice",
			child: []testCaseForBenchmark{
				{
					name: "in loop",
					child: []testCaseForBenchmark{
						{
							name: "int32",
							child: []testCaseForBenchmark{
								{
									name: "val",
									child: []testCaseForBenchmark{
										{
											name: "nzd",
											fvnc: benchSliceInLoopInt32,
										},
										{
											name: "zed",
											fvnc: benchSliceInLoopInt32Zed,
										},
									},
								},
								{
									name: "ptr",
									child: []testCaseForBenchmark{
										{
											name:  "nzd",
											fvnc:  benchSliceInLoopPtrInt32,
											alloc: 1,
											bytes: 4,
										},
										{
											name: "zed",
											fvnc: benchSliceInLoopPtrInt32Zed,
										},
									},
								},
							},
						},
					},
				},
				{
					name: "outside of loop",
					child: []testCaseForBenchmark{
						{
							name: "int32",
							child: []testCaseForBenchmark{
								{
									name: "val",
									child: []testCaseForBenchmark{
										{
											name: "nzd",
											fvnc: benchSliceOutLoopInt32,
										},
										{
											name: "zed",
											fvnc: benchSliceOutLoopInt32Zed,
										},
									},
								},
								{
									name: "ptr",
									child: []testCaseForBenchmark{
										{
											name:  "nzd",
											fvnc:  benchSliceOutLoopPtrInt32,
											alloc: 1,
											bytes: 4,
										},
										{
											name: "zed",
											fvnc: benchSliceOutLoopPtrInt32Zed,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	runTestCaseForBenchmark(t, testCases)
}

var (
	_leakDst *int32
	_moveDst *int32
)

func leak2Global(leaked2global *int32) *int32 {
	_leakDst = leaked2global
	return leaked2global
}

func leak2Result(leaked2result *int32) *int32 {
	return leaked2result
}

type testCaseForBenchmark struct {
	name  string
	fvnc  func(*testing.B)
	alloc int64
	bytes int64
	child []testCaseForBenchmark
}

func eqBytes(exp, act int64) bool {
	if exp == fifteenMiB {
		return (exp / oneMiB) == (act / oneMiB)
	}
	return exp == act
}

func runTestCaseForBenchmark(t *testing.T, testCases []testCaseForBenchmark) {
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			if tc.fvnc != nil {
				r := testing.Benchmark(tc.fvnc)
				if ea, aa := tc.alloc, r.AllocsPerOp(); ea != aa {
					t.Errorf("exp.alloc=%d, act.alloc=%d", ea, aa)
				}
				if eb, ab := tc.bytes, r.AllocedBytesPerOp(); !eqBytes(eb, ab) {
					t.Errorf("exp.bytes=%d, act.bytes=%d", eb, ab)
				}
				if !t.Failed() {
					t.Logf("allocs=%d, bytes=%d", tc.alloc, tc.bytes)
				}
			}
			runTestCaseForBenchmark(t, tc.child)
		})
	}
}

func benchPkgPointerToStackVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var stackVar int32 = 4096
		_moveDst = &stackVar
	}
}

func benchTooLargeForStackFrame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tooLargeForStackFrame [fifteenMiB]byte
		_ = tooLargeForStackFrame
	}
}

func benchNoEscape(b *testing.B) {
	noop := func(args ...interface{}) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var (
			a = new(int32)
			b = new(int64)
			c = make([]int32, 5, 5)
			d = make([]*int64, 10, 10)
			s = struct {
				a int32
				b int32
			}{a: 4096, b: 4096}
			ps = &struct {
				a int32
				b int32
			}{a: 4096, b: 4096}
		)
		noop(a, b, c, d, s, ps)
	}
}

func benchEscapeByLeakingToPkgVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := new(int32)
		*p = 4096
		leak2Global(p)
	}
}

func benchLeak2ResultStoreNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := new(int32) // Does not escape
		*p = 4096
		leak2Result(p)
	}
}

func benchLeak2ResultStoreInLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var pp *int32
		p := new(int32) // Does not escape
		*p = 4096
		pp = leak2Result(p)
		_ = pp
	}
}

func benchLeak2ResultStoreOutLoop(b *testing.B) {
	var pp *int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := new(int32) // Escapes
		*p = 4096
		pp = leak2Result(p)
	}
	b.StopTimer()
	_ = pp
}

func benchChanInt32Zed(b *testing.B) {
	c := make(chan int32, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32
		c <- x
		<-c
	}
}

func benchChanInt32(b *testing.B) {
	c := make(chan int32, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32 = 4096
		c <- x
		<-c
	}
}

func benchChanPtrInt32Zed(b *testing.B) {
	c := make(chan *int32, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p *int32
		c <- p
		<-c
	}
}

func benchChanPtrInt32(b *testing.B) {
	c := make(chan *int32, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := new(int32)
		c <- p
		<-c
	}
}

func benchIfaceInLoopInt32Zed(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst interface{}
		var x int32
		dst = x
		_ = dst
	}
}

func benchIfaceInLoopInt32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst interface{}
		var x int32 = 4096
		dst = x
		_ = dst
	}
}

func benchIfaceInLoopPtrInt32Zed(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst interface{}
		var p *int32
		dst = p
		_ = dst
	}
}

func benchIfaceInLoopPtrInt32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var dst interface{}
		p := new(int32)
		dst = p
		_ = dst
	}
}

func benchIfaceOutLoopInt32Zed(b *testing.B) {
	var dst interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32
		dst = x
	}
	_ = dst
}

func benchIfaceOutLoopInt32(b *testing.B) {
	var dst interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32 = 4096
		dst = x
	}
	_ = dst
}

func benchIfaceOutLoopPtrInt32Zed(b *testing.B) {
	var dst interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p *int32
		dst = p
	}
	_ = dst
}

func benchIfaceOutLoopPtrInt32(b *testing.B) {
	var dst interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := new(int32)
		dst = p
	}
	_ = dst
}

func benchSliceInLoopInt32Zed(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]int32, 1, 1)
		var x int32
		dst[0] = x
		_ = dst
	}
}

func benchSliceInLoopInt32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]int32, 1, 1)
		var x int32 = 4096
		dst[0] = x
		_ = dst
	}
}

func benchSliceInLoopPtrInt32Zed(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]*int32, 1, 1)
		var p *int32
		dst[0] = p
		_ = dst
	}
}

func benchSliceInLoopPtrInt32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]*int32, 1, 1)
		p := new(int32)
		dst[0] = p
		_ = dst
	}
}

func benchSliceOutLoopInt32Zed(b *testing.B) {
	dst := make([]int32, 1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32
		dst[0] = x
	}
	_ = dst
}

func benchSliceOutLoopInt32(b *testing.B) {
	dst := make([]int32, 1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32 = 4096
		dst[0] = x
	}
	_ = dst
}

func benchSliceOutLoopPtrInt32Zed(b *testing.B) {
	dst := make([]*int32, 1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p *int32
		dst[0] = p
	}
	_ = dst
}

func benchSliceOutLoopPtrInt32(b *testing.B) {
	dst := make([]*int32, 1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := new(int32)
		dst[0] = p
	}
	_ = dst
}
