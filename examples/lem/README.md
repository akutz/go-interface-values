# LEM (leak, escape, move)

The command below builds the program `./examples/lem/main.go` and illustrates the different scenarios where escape analysis is applied:

```bash
go build -gcflags "-l -N -m" ./examples/lem
```

The output should resemble the following:

```bash
# go-interface-values/examples/lem
examples/lem/main.go:22:2: leaking param: username
examples/lem/main.go:23:2: password does not escape
examples/lem/main.go:24:2: leaking param: token to result ~r0 level=0
examples/lem/main.go:30:10: "newLoginToken" escapes to heap
examples/lem/main.go:40:3: moved to heap: username1
examples/lem/main.go:43:3: moved to heap: cookieJar
examples/lem/main.go:41:18: new(string) escapes to heap
```
