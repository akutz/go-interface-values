# On the stack

We have learned that storing a value in an interface results in a copy of that value being created...somewhere. One of the possible locations is on the stack. To understand how a value is stored in an interface on the stack, we will take a look at the assembly. Please note:

* Assembly is not my area of expertise. Not even close. Therefore I am empathetic to people who may get stuck here. Please feel free to ping me on Gopher slack `@akutz` with any questions!
* This page tries to provide as many links and answer as many questions as possible regarding the assembly.
* Lastly, assembly is platform dependent. For example, the assembly for x86 does not look like the assembly for ARM. This page is based on x86 assembly.

The example on this page is based on the source code in [ifaceonthestack.go](../../examples/ifaceonthestack.go):

```go
/* line 17 */ package examples
/* line 18 */ 
/* line 19 */ func ifaceonthestack() {
/* line 20 */ 	var x int64
/* line 21 */ 	var y interface{}
/* line 22 */ 	x = 2
/* line 23 */ 	y = x
/* line 24 */ 	_ = y
/* line 25 */ }
```

With that in mind, let's get started:

1. Build the `examples` package with compiler flags to prevent write barriers (`-wb=false`), inlining (`-l`), and optimization (`-N`). You would never do this in producton, but it makes walking the assembly easier:

    ```bash
    go build -gcflags "-wb=false -l -N" -o examples.a ./examples
    ```

1. Dump the symbol `ifaceonthestack$` from the newly built archive:

    ```bash
    go tool objdump -s ifaceonthestack$ examples.a
    ```

---

:wave: **Alternative assembly** 

Please note it is also possible to dump the assembly for a single Go source file:

```bash
go tool compile -wb=false -l -N -S ./examples/ifaceonthestack.go
```

However I [have found](https://gophers.slack.com/archives/C029RQSEE/p1644033676178239) the Go compiler will produce different assembly based on `go tool compile` and actually packing the archive with `go build`. In order to be more aligned with package archive assembly, this page uses `go build`.

---

3. The resulting output will be similar (but not identical) to the following:

    ```assembly
    TEXT go-interface-values/examples.ifaceonthestack(SB) gofile../Users/akutz/Projects/go-interface-values/examples/ifaceonthestack.go
      ifaceonthestack.go:19	0x2c64			4883ec28		SUBQ $0x28, SP		[2:2]R_USEIFACE:type.int64	
      ifaceonthestack.go:19	0x2c68			48896c2420		MOVQ BP, 0x20(SP)	
      ifaceonthestack.go:19	0x2c6d			488d6c2420		LEAQ 0x20(SP), BP	
      ifaceonthestack.go:20	0x2c72			48c7042400000000	MOVQ $0x0, 0(SP)	
      ifaceonthestack.go:21	0x2c7a			440f117c2410		MOVUPS X15, 0x10(SP)	
      ifaceonthestack.go:22	0x2c80			48c7042402000000	MOVQ $0x2, 0(SP)	
      ifaceonthestack.go:23	0x2c88			48c744240802000000	MOVQ $0x2, 0x8(SP)	
      ifaceonthestack.go:23	0x2c91			488d0500000000		LEAQ 0(IP), AX		[3:7]R_PCREL:type.int64	
      ifaceonthestack.go:23	0x2c98			4889442410		MOVQ AX, 0x10(SP)	
      ifaceonthestack.go:23	0x2c9d			488d442408		LEAQ 0x8(SP), AX	
      ifaceonthestack.go:23	0x2ca2			4889442418		MOVQ AX, 0x18(SP)	
      ifaceonthestack.go:25	0x2ca7			488b6c2420		MOVQ 0x20(SP), BP	
      ifaceonthestack.go:25	0x2cac			4883c428		ADDQ $0x28, SP		
      ifaceonthestack.go:25	0x2cb0			c3			RET
    ```

    Here are the lines on which we want to focus:

    ```assembly
      ifaceonthestack.go:20	0x2c72			48c7042400000000	MOVQ $0x0, 0(SP)	
      ifaceonthestack.go:21	0x2c7a			440f117c2410		MOVUPS X15, 0x10(SP)	
      ifaceonthestack.go:22	0x2c80			48c7042402000000	MOVQ $0x2, 0(SP)	
      ifaceonthestack.go:23	0x2c88			48c744240802000000	MOVQ $0x2, 0x8(SP)	
      ifaceonthestack.go:23	0x2c91			488d0500000000		LEAQ 0(IP), AX		[3:7]R_PCREL:type.int64	
      ifaceonthestack.go:23	0x2c98			4889442410		MOVQ AX, 0x10(SP)	
      ifaceonthestack.go:23	0x2c9d			488d442408		LEAQ 0x8(SP), AX	
      ifaceonthestack.go:23	0x2ca2			4889442418		MOVQ AX, 0x18(SP)	
    ```

1. `ifaceonthestack.go:20	0x2c72			48c7042400000000	MOVQ $0x0, 0(SP)`
    * `ifaceonthestack.go:20`
        * This is the file and line number of the source code that corresponds to this line of assembly.
        * In this case it is line 20 from the file `ifaceonthestack.go` -- `var x int64`.
    * `0x2c72`
        * The program counter formatted as hexadecimal.
        * GNU's `objdump` tool formats this value as hexadecimal as well, but without the leading prefix `0x`.
    * `48c7042400000000`
        * The executable instruction formatted as hexadecimal.
        * GNU's `objdump` tool formats this value as hexadecimal as well, but with spaces, ex. `48 c7 04 24 00 00 00 00`.
    * `MOVQ $0x0, 0(SP)`
        * The instruction `MOVQ` copies the value from one address to another, ex. `MOVQ SRC, DST`.
        * `MOVQ`
            * Normally `MOVQ` operates right to left, `MOVQ DST SRC`, but as the [Go assembly documentation](https://go.dev/doc/asm) states:

                > One detail evident in the examples from the previous sections is that data in the instructions flows from left to right: MOVQ $0, CX clears CX. This rule applies even on architectures where the conventional notation uses the opposite direction. 
            * The `Q` in `MOVQ` stands for _quadword_:
                * On x86 and x86_64 platforms a _word_ is 16 bits.
                * On 64-bit platforms a _quadword_ is 16x4, or 64-bits.
                * Thus `MOVQ` is used when wanting to copy 8 bytes.
        * `$0x0`
            * The `SRC` of the copy operation.
            * The leading `$` indicates `SRC` is not a memory address, but a literal value.
            * The value to copy is therefore `0x0`, or the integer value `0`.
        * `0(SP)`
            * The `DST` of the copy operation.
            * The `0` indicates an offset of zero bytes from some address.
            * The address is indicated by `(SP)`, _stack pointer_, the highest memory address of the current stack frame.
            * Therefore `0(SP)` can be translated as _zero bytes from the highest memory address of the current strack frame_.

    <br />

    ```
    SP +---------------------------------------+ SP + 0 bytes
       | name: x                               |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 0                               |
       +---------------------------------------+
    ```

1. `ifaceonthestack.go:21	0x2c7a			440f117c2410		MOVUPS X15, 0x10(SP)`
    * The assembly for line21, `var y interface{}`.
    * `MOVUPS X15 0x10(SP)`
        * `MOVUPS`
            * The instruction [`MOVUPS`](https://www.felixcloutier.com/x86/movups) copies an unligned, packed, single-precision floating point value from one address to another, ex. `MOVUPS SRC, DST`.
            * Like `MOVQ`, when reading Go assembly `MOVUPS` operates right-to-left, `DST SRC`.
            * Go is using `MOVUPS` in 128-bit mode, which means the operation is copying 16 bytes.
        * `X15`
            * The `SRC` of the copy operation.
            * The `X15` register is special and holds the zero value (Go application binary interface (ABI) [documentation](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md)).
            * Because `MOVUPS` is copying 16 bytes of data and the `X15` register is `0`, this instruction is essentially reserving 16 bytes on the stack starting at `DST`.
        * `0x10(SP)`
            * The `DST` of the copy operation.
            * The `0x10`  indicates an offset of 16 bytes (`0x10` is hexadecimal for 16) from some address.
            * The address is indicated by `(SP)`, _stack pointer_, the highest memory address of the current stack frame.
            * Therefore `0x10(SP)` can be translated as _16 bytes from the highest memory address of the current strack frame_.

    <br />

    ```
    SP +---------------------------------------+ SP + 0 bytes
       | name: x                               |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 0                               |
       +---------------------------------------+ SP + 8 bytes
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       +---------------------------------------+ SP + 16 bytes
       | name: y                               |
       | type: interface{}                     |
       | size: 16 bytes                        |
       | valu: <reserved>                      |
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       +---------------------------------------+
    ```

    Wait, why was `y` offset by 16 bytes when `x` is only eight bytes? Find out below! :smiley:

1. `ifaceonthestack.go:22	0x2c80			48c7042402000000	MOVQ $0x2, 0(SP)`
    * The assembly for `x = 2`
    * `MOVQ $0x2, 0x8(SP)` copies the literal value `2` to the memory address for the variable `x`.

    <br />

    ```
    SP +---------------------------------------+ SP + 0 bytes
       | name: x                               |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 2                               |
       +---------------------------------------+ SP + 8 bytes
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       +---------------------------------------+ SP + 16 bytes
       | name: y                               |
       | type: interface{}                     |
       | size: 16 bytes                        |
       | valu: <reserved>                      |
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       +---------------------------------------+
    ```

1. `ifaceonthestack.go:23	0x2c88			48c744240802000000	MOVQ $0x2, 0x8(SP)`
    * The assembly for `y = x`
    * `MOVQ $0x2, 0x10(SP)` copies the literal value `2` to the memory address 16 bytes from `SP`.
    * Please note this is not a named variable, or rather not a named `int64`.
    * The Go compiler was able to determine that the only value ever assigned to `y` would be an `int64`, and so an extra eight bytes was allocated on the stack in order to store the `int64` value assigned to `y`.

    <br />

    ```
    SP +--------------------------------------+ SP + 0 bytes
       | name: x                              |
       | type: int64                          |
       | size: 8 bytes                        |
       | valu: 2                              |
       +--------------------------------------+ SP + 8 bytes
       | name:                                |
       | type: int64                          |
       | size: 8 bytes                        |
       | valu: 2                              |
       +--------------------------------------+ SP + 16 bytes
       | name: y                              |
       | type: interface{}                    |
       | size: 16 bytes                       |
       | valu: <reserved>                     |
       |                                      |
       |                                      |
       |                                      |
       |                                      |
       |                                      |
       +--------------------------------------+
    ```

1. `ifaceonthestack.go:23	0x2c91			488d0500000000		LEAQ 0(IP), AX		[3:7]R_PCREL:type.int64`
    * Still more assembly for `y = x`
    * `LEAQ`
        * The [`LEA`](https://www.felixcloutier.com/x86/lea) instruction stands for _load effective address_.
        * The `Q` suffix indicates a _quadword_, aka 64 bits, aka 8 bytes.
        * Like other Go assembly, the `DST SRC` syntax is flipped to be `SRC DST`
        * Unlike `MOV` which reads the memory at the provided `SRC` address, `LEA` only reads the address itself. For example, the code snippet below would result in a `MOV` instruction in order to copy the value of `x` (address `0(SP)`) to the address of `y` (address `0x8(SP)`):

            ```go
            x := 1  // MOVQ $0x1  0(SP)
            y := x  // MOVQ 0(SP) 0x8(SP)
            ```

            Actually, the Go compiler is pretty smart, and it would probably use `MOVQ $0x1 0x8(SP)` to assign `1` to `y`, but for the purposes of this example we copied the value of `x` to `y`. However, _this_ code snippet would use an `LEA` since we do not need to know the value of `x`, only its address:

            ```go
            x := 1  // MOVQ $0x1  0(SP)
            y := &x // LEAQ 0(SP) 0x8(SP)
            ```

    * `LEAQ 0(IP), AX` stores the address of the instruction `[3:7]R_PCREL:type.int64` in register `AX`.
    * The symbol [`R_PCREL`](https://developer.apple.com/documentation/kernel/scattered_relocation_info/1577780-r_pcrel) is specific to darwin and indicates the item containing the instruction uses program counter relative addressing.
    * Ultimately what is stored in `AX` is the address of `type.int64`, a global value that specifies the internal type for an `int64`.

1. `ifaceonthestack.go:23	0x2c88			48c744240802000000	MOVQ $0x2, 0x8(SP)`
    * Still more assembly for `y = x`
    * `MOVQ AX, 0x10(SP)` copies the value in the `AX` register to the memory address offset from `SP` by 16 bytes.
    * This assigns the address of the global value `type.int64` to the interface's first `uintptr`, the one that points to the underlying type.

    <br />

    ```
    SP +---------------------------------------+ SP + 0 bytes
       | name: x                               |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 2                               |
       +---------------------------------------+ SP + 8 bytes
       | name:                                 |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 2                               |
       +-------------------+-------------------+ SP + 16 bytes
       | name: y           | type: uintptr     |
       | type: interface{} | size: 8 bytes     |
       | size: 16 bytes    | valu: &type.int64 |
       |                   |                   |
       |                   +-------------------+
       |                                       |
       |                                       |
       |                                       |
       |                                       |
       +---------------------------------------+
    ```

1. `ifaceonthestack.go:23	0x2c9d			488d442408		LEAQ 0x8(SP), AX`
    * Still more assembly for `y = x`
    * `LEAQ 0x8(SP), AX` loads the address of the memory eight bytes from `SP` into the register `AX`.
    * The address loaded into `AX` points to the aforementioned, unnamed, temporary value the Go compiler created on the stack for the `y` interface to reference.

1. `ifaceonthestack.go:23	0x2ca2			4889442418		MOVQ AX, 0x18(SP)`
    * Still more assembly for `y = x`
    * `MOVQ AX, 0x18(SP)` copies the value in register `AX` into the address 24 bytes from `SP`.
    * This assigns the address of the unamed `int65` at `0x8(SP)` to the interface's second `uintptr`, the one that points to the underlying value.

    <br />

    ```
    SP +---------------------------------------+ SP + 0 bytes
       | name: x                               |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 2                               |
       +---------------------------------------+ SP + 8 bytes
       | name:                                 |
       | type: int64                           |
       | size: 8 bytes                         |
       | valu: 2                               |
       +-------------------+-------------------+ SP + 16 bytes
       | name: y           | type: uintptr     |
       | type: interface{} | size: 8 bytes     |
       | size: 16 bytes    | valu: &type.int64 |
       |                   |                   |
       |                   +-------------------+ SP + 24 bytes
       |                   | type: uintptr     |
       |                   | size: 8 bytes     |
       |                   | valu: 0x8(SP)     |
       |                   |                   |
       +-------------------+-------------------+
    ```

Wait a minute, isn't memory referenced by pointer allocated on the heap!? In fact Go can optimize that memory to the stack as well, and that is what happens in this example. The Go compiler was able to place the value stored in `y` on the stack at address `0x8(SP)` and let the pointer at `0x18(SP)` reference `0x8(SP)`, all on the stack.

However, the fact that a "temporary" memory location was created to reference from the interface is key to understanding when storing an interface value results in a memory allocation on the heap. Still, before we answer _that_ question, we first need to understand why the Go compiler was able to keep _this_ value on the stack. 

Keep reading to learn about escape analysis!

---

Next: [Escape analysis](../03-escape-analysis/)