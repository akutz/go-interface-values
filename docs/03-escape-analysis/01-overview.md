# Overview

Go reclaims memory allocated on the heap via garbage collection, but this can be a very expensive process. Compare that to the stack where memory is "cheap" and is freed automatically when its stack frame is destroyed. In order to allocate memory on the stack the Go compiler must evaluate several, determining factors:

* pointers to stack objects cannot be stored in the heap
* pointers to stack objects cannot outlive the object's stack frame
* stack objects cannot exceed the size of the stack, ex. a 15 MiB buffer `[15 * 1024 * 1024]byte`

The compile-time process Go use to determine whether memory is dynamically managed on the heap or can be allocated on the stack is known as [_escape analysis_](https://github.com/golang/go/blob/master/src/cmd/compile/internal/escape/escape.go). Escape analysis walks a program's [abstract syntax tree](https://pkg.go.dev/go/ast) (AST) to build a graph of all the variables encountered.

It is possible to see which variables end up on the heap by using the compiler flag `-m` when building (or testing) Go code. For example, let's build the program at [`./examples/cmd/lem/main.go`](../../examples/cmd/lem/main.go) with the the following command:

```bash
docker run -it --rm go-interface-values go build -gcflags "-m" ./examples/cmd/lem
```

The output should be:

```
# go-interface-values/examples/cmd/lem
examples/cmd/lem/main.go:25:18: leaking param: idToTrack
examples/cmd/lem/main.go:30:17: leaking param: idToValidate to result ~r0 level=0
examples/cmd/lem/main.go:41:3: moved to heap: userData
examples/cmd/lem/main.go:43:2: userID escapes to heap
```

In one, very small bit of code we have variables that are:

* leaked to the heap
* escaping to the heap
* and moved to the heap

Please keep reading to find out when variables leak, escape, and are moved to the heap!

---

Next: [Leak](./02-leak.md)
