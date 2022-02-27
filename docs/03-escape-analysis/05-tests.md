# Tests

This section on escape analysis has been rather assertive in its claims about criteria and behaviors. Truthfully this entire repository is the embodiment of "one physical test is worth one thousand expert opinions." I am no expert on, well, much of anything, but I do like digging into problems, even those I do not fully understand. Storing interface values, heap allocation, and escape analysis are such problems. The sources are there in Go, and I can follow _just_ enough to be annoying on Gopher Slack :smiley:

To that end this repository has several tests to verify the assertions regarding escape analysis. The tests are located in [`./tests/lem`](../../tests/lem) and may be executed with the following command:

```bash
docker run -it --rm go-interface-values go test -v ./tests/lem
```

The tests use a bespoke framework called _leak, escape, move_, or "lem" for short. Lem enables the use of code comments decorating and within functions to assert the expected number of allocations, bytes allocated, and where the compiler's optimiation output should indicate a value has leaked, escaped, or moved to the heap. For example:

```go
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
```

It is even possible to write tests that assert when optimizations should _not_ occur:

```go
// lem.leak7.name=did not escape bc return value did not outlive stack frame
// lem.leak7.alloc=0
// lem.leak7.bytes=0
func leak7(b *testing.B) {
	f := func(p *int32) *int32 { // lem.leak7.m=leaking param: p to result ~r[0-1] level=0
		return p
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var x int32 = 4096 // lem.leak7.m!=(escapes|leaking|moved)
		var p *int32 = &x  // lem.leak7.m!=(escapes|leaking|moved)
		var sink *int32    // lem.leak7.m!=(escapes|leaking|moved)
		sink = f(p)        // lem.leak7.m!=(escapes|leaking|moved)
		_ = sink
	}
}
```

For more information on the lem test framework please refer [github.com/akutz/lem](https://github.com/akutz/lem). Go ahead, I'll wait. No, seriously, I'll be right here.

...

You're back? Great! So now that we all understand the basics of escape analysis it is time to revisit why storing values in interfaces does not always result in memory allocated on the heap.

---

Next: [Missing mallocs](../04-missing-mallocs/)
