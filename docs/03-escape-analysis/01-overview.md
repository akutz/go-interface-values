# Overview

Go reclaims memory allocated on the heap via garbage collection, but this can be a very expensive process. Compare that to the [stack](https://github.com/golang/go/blob/master/src/runtime/stack.go) where memory is "cheap" and is freed automatically when its stack frame is destroyed. In order to allocate memory on the stack the Go compiler must evaluate several, determining factors:

* pointers to stack objects cannot be stored in the heap
* pointers to stack objects cannot outlive the object's stack frame
* stack objects cannot exceed the size of the stack, ex. a 15 MiB buffer `[15 * 1024 * 1024]byte`

The compile-time process Go use to determine whether memory is dynamically managed on the heap or can be allocated on the stack is known as [_escape analysis_](https://github.com/golang/go/blob/master/src/cmd/compile/internal/escape/escape.go). Escape analysis walks a program's [abstract syntax tree](https://pkg.go.dev/go/ast) (AST) to build a graph of all the variables encountered.

It is possible to see which variables end up on the heap by using the compiler flag `-m` when building (or testing) Go code. For example, let's build the program at [`./examples/lem/main.go`](../../examples/lem/main.go):

```bash
docker run -it --rm go-interface-values go build -gcflags "-l -m -N" ./examples/lem
```

The output should be similar to the following:

```
# go-interface-values/examples/lem
examples/lem/main.go:22:2: leaking param: username
examples/lem/main.go:23:2: password does not escape
examples/lem/main.go:24:2: leaking param: token to result ~r0 level=0
examples/lem/main.go:30:10: "newLoginToken" escapes to heap
examples/lem/main.go:40:3: moved to heap: username1
examples/lem/main.go:43:3: moved to heap: cookieJar
examples/lem/main.go:41:18: new(string) escapes to heap
```

This tiny program has variables that _leak_, _escape_, and get _moved_ to the heap. Please keep reading to find out what these terms mean and when and why they occur.

---

Next: [Leak](./02-leak.md)
