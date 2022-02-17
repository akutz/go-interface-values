# Lessons learned

Wow, what a ride! What I meant to be a quick detour to help me better understand what performance improvements generics might bring compared to storing values in interfaces turned into quite a deep-dive. What did we learn along the way?

* While storing a value in an interface is conceptually similar to "boxing," this term has a specific meaning for object oriented languages and may confuse people.
* A Go interface is really just a pair of `uintptr` values that point to the interface's underlying type and value.
* Escape analysis occurs before optimizations for storing values in an interface.
* Just because a value is marked as "escapes to heap" does not mean memory is allocated on the heap.

And remember, if you ever want to see if an assignment allocates memory on the heap, all you have to do is take a quick trip to the [Go playground](https://go.dev/play/p/OoN-qxRHqsi):

```go
package main

import (
	"flag"
	"testing"
)

func TestStoringInterfaceValues(t *testing.T) {
	flag.Lookup("test.benchmem").Value.Set("true")
	flag.Lookup("test.benchtime").Value.Set("100x")

	testCases := []struct {
		name     string
		fvnc     func(*testing.B)
		expAlloc int64
		expBytes int64
	}{
		{
			name: "0 malloc, 0 bytes",
			fvnc: func(b *testing.B) {
				var sink interface{}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var x int32
					x = 255
					sink = x
				}
				_ = sink
			},
		},
		{
			name: "1 malloc, 4 bytes",
			fvnc: func(b *testing.B) {
				var sink interface{}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var x int32
					x = 256
					sink = x
				}
				_ = sink
			},
			expAlloc: 1,
			expBytes: 4,
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			r := testing.Benchmark(tc.fvnc)
			if ea, aa := tc.expAlloc, r.AllocsPerOp(); ea != aa {
				t.Errorf("expAlloc=%d, actAlloc=%d", ea, aa)
			}
			if eb, ab := tc.expBytes, r.AllocedBytesPerOp(); eb != ab {
				t.Errorf("expBytes=%d, actBytes=%d", eb, ab)
			}
		})
	}
}
```

```bash
=== RUN   TestStoringInterfaceValues
=== RUN   TestStoringInterfaceValues/0_malloc,_0_bytes
=== RUN   TestStoringInterfaceValues/1_malloc,_4_bytes
--- PASS: TestStoringInterfaceValues (0.00s)
    --- PASS: TestStoringInterfaceValues/0_malloc,_0_bytes (0.00s)
    --- PASS: TestStoringInterfaceValues/1_malloc,_4_bytes (0.00s)
PASS
```

---

_Thanks for reading!_
