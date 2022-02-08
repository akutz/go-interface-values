# Golang interface values and memory

This repository explores when storing a value in a Go interface allocates memory on the heap.

* [**Labs**](#labs): a step-by-step walkthrough of the topic
* [**FAQ**](#FAQ): answers to frequently asked questions
* [**Links**](#links): links to related reference material

## Labs

1. [**Prerequisites**](./docs/01-prereqs/): how to get from here to there
1. [**Interface values**](./docs/02-interface-values/): whatever you do, do not call it "boxing"
1. [**Escape analysis**](./docs/03-escape-analysis/): to malloc or not to malloc
1. [**Missing mallocs**](./docs/04-missing-mallocs/): there's a heap of missing memory

## FAQ

* [**What is the x86 assembly instruction `CALL` actually calling?**](#what-is-the-x86-assembly-instruction-actually-calling)
* [**Where is the `CALL` instruction in ARM assembly?**](#where-is-the-call-instruction-in-arm-assembly)
* [**What is the `hack` directory and the files inside of it?**](#what-is-the-hack-directory-and-the-files-inside-of-it)

### What is the x86 assembly instruction `CALL` actually calling?

The x86 assembly instruction `CALL` is used to call a procedure. For example, consider the following line of assembly from this project:

```assembly
0x003c 00060 (mem_test.go:312)	CALL	runtime.convT32(SB)
```

The `CALL` instruction specifies `runtime.convT32(SB)`. The `(SB)` (_static base_) suffix lets us know that `runtime.convT32` is a global symbol and references the memory address for the procedure to call. In order to examine what that symbol is, we can:

1. Build the test binary for this project:

    ```bash
    go test -c ./benchmarks
    ```

2. Use `go tool objdump` to search for the symbol:

    ```bash
    $ go tool objdump -s runtime.convT32 benchmarks.test
    TEXT runtime.convT32(SB) /Users/akutz/.go/1.18beta2/src/runtime/iface.go
      iface.go:365		0x100008b20		f9400b90		MOVD 16(R28), R16
      ...
    ```

    Procedures in assembly are easy to spot as they are defined with the `TEXT` directive illustrated up above.
    
3. Or build the `runtime` package without linking it into this project's test binary. This assembly provides even more insight into the translation from Go to machine code:

    ```bash
    $ go build -gcflags "-S" runtime 2>&1 | grep -FA1 '"".convT32 STEXT'
    "".convT32 STEXT size=144 args=0x8 locals=0x28 funcid=0x0 align=0x0
    	0x0000 00000 (/Users/akutz/.go/1.18beta2/src/runtime/iface.go:365)	TEXT	"".convT32(SB), ABIInternal, $48-8
    	...
    ```

To learn more about the `CALL` instruction and procedures, please refer to:

* [**A Quick Guide to Go's Assembler**](https://go.dev/doc/asm) ([symbols](https://go.dev/doc/asm#symbols), [directives](https://go.dev/doc/asm#directives))
* [**x86 asm `CALL`**](https://www.felixcloutier.com/x86/call)

### Where is the `CALL` instruction in ARM assembly?

The assembly for x86 and amd64 both define the [`CALL`](https://www.felixcloutier.com/x86/call) instruction. However, [ARM assembly](https://developer.arm.com/documentation/ddi0602/2021-12/?lang=en) does _not_ define a `CALL` instruction, despite it appearing in the assembly for Go sources in this project when they are compiled on an M1 Macbook. What gives?

What was likely an attempt to maintain naming convention consistency in Go assembly across different processor architectures, Go renames the ARM assembly instruction [`BLR`](https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/BLR--Branch-with-Link-to-Register-?lang=en) (_branch with link to register_) to `CALL`. Both instructions call a procedure (x86) or subroutine (ARM) at a given address, so for all intents and purposes they have semantically similar.

To learn more about the `CALL` instruction and Go on ARM, please refer to:

* [**Instructions mnemonics mapping rules for the Go ARM64 assembler**](https://pkg.go.dev/cmd/internal/obj/arm64#hdr-Instructions_mnemonics_mapping_rules)
* [**ARM64 developer documentation**](https://developer.arm.com/documentation/ddi0602/2021-12/?lang=en)

### What is the [`hack`](./hack) directory and the files inside of it?

The `hack` directory is a convention I picked up from working on Kubernetes and projects related to Kuberentes. The directory contains scripts useful to the project, but not a core piece of the project itself. For example:

* [**`hack/`**](./hack)
  * [**`asm2md.py`**](./hack/asm2md.py): parses the output of `go tool compile -S -wb=false *.go` and produces a markdown table
  * [**`b2md.py`**](./hack/b2md.py): parses the output of `go test -bench BenchmarkMem -run Mem -benchmem -count 1 -benchtime 1000x -v` and produces a markdown table 
  * [**`gen.py`**](./hack/gen.py): generates [`mem_test.go`](mem_test.go), [`print_test.go`](print_test.go), and [`types_test.go`](types_test.go)

## Links

* [**ARM developer documentation**](https://developer.arm.com/documentation/ddi0602/2021-12/?lang=en)
* [**x86 and amd64 instruction set**](https://www.felixcloutier.com/x86/index.html)
* [**A quick guide to Go assembly**](https://go.dev/doc/asm)
