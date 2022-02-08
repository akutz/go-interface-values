# Overview

Go tries really, _really_ hard to keep memory on the stack where possible. Memory allocated on the stack relieves pressure on the garbage collector as the memory is cleaned up once the stack on which it is allocated no longer exists.

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
