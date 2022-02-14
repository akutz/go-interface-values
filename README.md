# Go interface values

This repository deep dives Go interface values, what they are, how they work, and when storing a value in a Go interface allocates memory on the heap.

* [**Labs**](#labs): a step-by-step walkthrough of the topic
* [**FAQ**](#FAQ): answers to frequently asked questions
* [**Links**](#links): links to related reference material
* [**Appendix**](#appendix): in-repo reference material


## Labs

1. [**Prerequisites**](./docs/01-prereqs/): how to get from here to there
1. [**Interface values**](./docs/02-interface-values/): whatever you do, do not call it "boxing"
1. [**Escape analysis**](./docs/03-escape-analysis/): to malloc or not to malloc
1. [**Missing mallocs**](./docs/04-missing-mallocs/): there's a heap of missing memory


## FAQ

* [**What does the `Q` suffix for instructions like `MOVQ` and `LEAQ` mean?**](#what-does-the-q-suffix-for-instructions-like-movq-and-leaq-mean)
* [**What is the x86 assembly instruction `CALL` actually calling?**](#what-is-the-x86-assembly-instruction-call-actually-calling)
* [**Where is the `CALL` instruction in ARM assembly?**](#where-is-the-call-instruction-in-arm-assembly)
* [**What is the `hack` directory and the files inside of it?**](#what-is-the-hack-directory-and-the-files-inside-of-it)


### What does the `Q` suffix for instructions like `MOVQ` and `LEAQ` mean?

Please refer to [this answer](./docs/99-appendix/assembly.md#what-does-the-q-suffix-for-instructions-like-movq-and-leaq-mean) from the assembly section in the appendix.


### What is the x86 assembly instruction `CALL` actually calling?

Please refer to [this answer](./docs/99-appendix/assembly.md#what-is-the-x86-assembly-instruction-actually-calling) from the assembly section in the appendix.


### Where is the `CALL` instruction in ARM assembly?

Please refer to [this answer](./docs/99-appendix/assembly.md#where-is-the-call-instruction-in-arm-assembly) from the assembly section in the appendix.


### What is the [`hack`](./hack) directory and the files inside of it?

The `hack` directory is a convention I picked up from working on Kubernetes and projects related to Kuberentes. The directory contains scripts useful to the project, but not a core piece of the project itself. For example:

* [**`hack/`**](./hack)
  * [**`asm2md.py`**](./hack/asm2md.py): parses the output of `go tool compile -S -wb=false *.go` and produces a markdown table
  * [**`b2md.py`**](./hack/b2md.py): parses the output of `go test -bench BenchmarkMem -run Mem -benchmem -count 1 -benchtime 1000x -v` and produces a markdown table
  * [**`gen.py`**](./hack/gen.py): generates [`./benchmarks/mem_test.go`](./benchmarks/mem_test.go), [`./benchmarks/print_test.go`](./benchmarks/print_test.go), and [`./benchmarks/types_test.go`](./benchmarks/types_test.go)


## Links

* [**ARM developer documentation**](https://developer.arm.com/documentation/ddi0602/2021-12/?lang=en)
* [**x86 and amd64 instruction set**](https://www.felixcloutier.com/x86/index.html)
* [**A quick guide to Go assembly**](https://go.dev/doc/asm)
* [**Logging, interfaces, and allocation**](https://commaok.xyz/post/interface-allocs/)
* [**Go introduction to escape analysise**](https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890)
* [**Type definitions for leaks**](https://github.com/golang/go/blob/master/src/cmd/compile/internal/escape/graph.go)
* [**Enums for escape**](https://github.com/golang/go/blob/master/src/cmd/compile/internal/ir/node.go)
* [**Type definition for a Go stack frame**](https://github.com/golang/go/blob/master/src/runtime/stack.go)


## Appendix

* [**Assembly**](./docs/99-appendix/assembly.md): reference section for go asm
