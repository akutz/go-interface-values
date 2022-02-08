# Assembly

This is a reference for the assembly found in this repository:

* [**Why are all of the operands inverted for `MOV` and other instructions?**](#why-are-all-of-the-operands-inverted-for-mov-and-other-instructions)
* [**What does the `Q` suffix for instructions like `MOVQ` and `LEAQ` mean?**](#what-does-the-q-suffix-for-instructions-like-movq-and-leaq-mean)
* [**What is the x86 assembly instruction `CALL` actually calling?**](#what-is-the-x86-assembly-instruction-actually-calling)
* [**Where is the `CALL` instruction in ARM assembly?**](#where-is-the-call-instruction-in-arm-assembly)


## Why are all of the operands inverted for `MOV` and other instructions?

Normally `MOVQ` operates right to left, `MOVQ DST SRC`, but as the [Go assembly documentation](https://go.dev/doc/asm) states:

> One detail evident in the examples from the previous sections is that data in the instructions flows from left to right: MOVQ $0, CX clears CX. This rule applies even on architectures where the conventional notation uses the opposite direction. 

This inversion is true for many different instructions with the `DST SRC` operands. In Go asm the order is inverted to be `SRC DST`.


## What does the `Q` suffix for instructions like `MOVQ` and `LEAQ` mean?

While a _word_ is supposed to be dependent upon the platform, because _word_ was eight bits in early x86 architecture, a _word_ still refers to eight bits on modern x86 and x86_64 systems. Therefore a _quadword_, or `Q`, refers to 64 bits.

Thus instrucitons like `MOVQ` and `LEAQ` are variants of `MOV` and `LEA` that address 64-bit literals and/or memory addresses. For example, in Go asm the instruction `MOVQ 0x8(SP) 0x10(SP)` copies eight bytes of data from the address `0x8(SP)` to `0x10(SP)`.


## What is the x86 assembly instruction `CALL` actually calling?

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


## Where is the `CALL` instruction in ARM assembly?

The assembly for x86 and amd64 both define the [`CALL`](https://www.felixcloutier.com/x86/call) instruction. However, [ARM assembly](https://developer.arm.com/documentation/ddi0602/2021-12/?lang=en) does _not_ define a `CALL` instruction, despite it appearing in the assembly for Go sources in this project when they are compiled on an M1 Macbook. What gives?

What was likely an attempt to maintain naming convention consistency in Go assembly across different processor architectures, Go renames the ARM assembly instruction [`BLR`](https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/BLR--Branch-with-Link-to-Register-?lang=en) (_branch with link to register_) to `CALL`. Both instructions call a procedure (x86) or subroutine (ARM) at a given address, so for all intents and purposes they have semantically similar.

To learn more about the `CALL` instruction and Go on ARM, please refer to:

* [**Instructions mnemonics mapping rules for the Go ARM64 assembler**](https://pkg.go.dev/cmd/internal/obj/arm64#hdr-Instructions_mnemonics_mapping_rules)
* [**ARM64 developer documentation**](https://developer.arm.com/documentation/ddi0602/2021-12/?lang=en)
