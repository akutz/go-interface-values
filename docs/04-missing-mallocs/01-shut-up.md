# Shut up and prove it

So how do we demonstrate that a value escaping to the heap does not result in a heap allocation? With a Go benchmark of course! Consider the following:

```go
var sink interface{}

func ex1() {
	var x int64
	x = 3
	sink = x
}
```

Based on all that we have learned, we can be reasonably certain when it is stored in `sink`, the value of `x` will escape to the heap. Let's find out!

```bash
docker run -it --rm go-interface-values \
  go tool compile -l -m ./docs/04-missing-mallocs/examples/ex1/ex1.go
```

The output should look similar to this:

```bash
./docs/04-missing-mallocs/examples/ex1/ex1.go:25:2: x escapes to heap
```

Sure enough, `x` escapes to the heap! But does it result in a heap allocation, that is the question! To find out we can a benchmark that invokes the above `ex1` function and track the allocations/bytes per operation:

```bash
docker run -it --rm go-interface-values \
  go test -gcflags -l -bench . -benchmem -v ./docs/04-missing-mallocs/examples/ex1/
```

The output will vary, but it is one line in particular from below that matters:

```bash
goos: linux
goarch: amd64
pkg: go-interface-values/docs/04-missing-mallocs/examples/ex1
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkEscapeNoMalloc
BenchmarkEscapeNoMalloc-8   	663767829	         1.846 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	go-interface-values/docs/04-missing-mallocs/examples/ex1	1.421s
```

That's right, this one:

```bash
BenchmarkEscapeNoMalloc-8   	663767829	         1.846 ns/op	       0 B/op	       0 allocs/op
```

There were _zero_ allocations and _zero_ bytes allocated on the heap. But how can that be? We have learned that:

* storing a value in an interface results in a copy of that value
* `x` escapes to the heap because its value outlives its stack frame by being assigned to the `sink`

This _should_ mean a new `int64` is allocated on the heap to store the copy of `x`, right? To find out what is happening we need to look a little closer...

---

Next: [Looking at the assembly](./02-why.md)
