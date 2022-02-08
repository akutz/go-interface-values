# LEM (leak, escape, move)

Building this program with the "-m" compiler flag will illustrate all three manners by which a variable can end up on the heap:

```bash
go build -gcflags "-m" ./examples/cmd/lem
```

The output of the above command should be identical to the following:

```bash
# go-interface-values/examples/cmd/lem
examples/cmd/lem/main.go:25:18: leaking param: idToTrack
examples/cmd/lem/main.go:30:17: leaking param: idToValidate to result ~r0 level=0
examples/cmd/lem/main.go:41:3: moved to heap: userData
examples/cmd/lem/main.go:43:2: userID escapes to heap
```
