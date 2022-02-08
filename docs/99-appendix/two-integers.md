# Two integers

This page describes how two integer values are allocated on the stack by walking assembly language. The example on this page is based on the source code in [xint2yint.go](../../examples/xint2yint.go):

```go
/* line 17 */ package examples
/* line 18 */ 
/* line 19 */ func xint2yint() {
/* line 20 */ 	var x int64
/* line 21 */ 	var y int64
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

1. Dump the symbol `xint2yint$` from the newly built archive:

    ```bash
    go tool objdump -s xint2yint$ examples.a
    ```

---

:wave: **Alternative assembly** 

Please note it is also possible to dump the assembly for a single Go source file:

```bash
go tool compile -wb=false -l -N -S ./examples/xint2yint.go
```

However I [have found](https://gophers.slack.com/archives/C029RQSEE/p1644033676178239) the Go compiler will produce different assembly based on `go tool compile` and actually packing the archive with `go build`. In order to be more aligned with package archive assembly, this page uses `go build`.

---

3. The resulting output depends on the platform. The following was produced on darwin/amd64:

    ```assembly
    TEXT go-interface-values/examples.xint2yint(SB) gofile../Users/akutz/Projects/go-interface-values/examples/xint2yint.go
      xint2yint.go:19	0x2c2a			4883ec18		SUBQ $0x18, SP		
      xint2yint.go:19	0x2c2e			48896c2410		MOVQ BP, 0x10(SP)	
      xint2yint.go:19	0x2c33			488d6c2410		LEAQ 0x10(SP), BP	
      xint2yint.go:20	0x2c38			48c744240800000000	MOVQ $0x0, 0x8(SP)	
      xint2yint.go:21	0x2c41			48c7042400000000	MOVQ $0x0, 0(SP)	
      xint2yint.go:22	0x2c49			48c744240802000000	MOVQ $0x2, 0x8(SP)	
      xint2yint.go:23	0x2c52			48c7042402000000	MOVQ $0x2, 0(SP)	
      xint2yint.go:25	0x2c5a			488b6c2410		MOVQ 0x10(SP), BP	
      xint2yint.go:25	0x2c5f			4883c418		ADDQ $0x18, SP		
      xint2yint.go:25	0x2c63			c3			RET
    ```

    We want to focus on four lines in particular:

    ```assembly
      xint2yint.go:20	0x2c38			48c744240800000000	MOVQ $0x0, 0x8(SP)	
      xint2yint.go:21	0x2c41			48c7042400000000	MOVQ $0x0, 0(SP)	
      xint2yint.go:22	0x2c49			48c744240802000000	MOVQ $0x2, 0x8(SP)	
      xint2yint.go:23	0x2c52			48c7042402000000	MOVQ $0x2, 0(SP)	
    ```

    Let's take it line by line.

1. `xint2yint.go:20	0x2c38			48c744240800000000	MOVQ $0x0, 0x8(SP)`
    * `xint2yint.go:20`
        * This is the file and line number of the source code that corresponds to this line of assembly.
        * In this case it is line 20 from the file `xint2yint.go` -- `var x int64`.
    * `0x2c38`
        * The program counter formatted as hexadecimal.
        * GNU's `objdump` tool formats this value as hexadecimal as well, but without the leading prefix `0x`.
    * `48c744240800000000`
        * The executable instruction formatted as hexadecimal.
        * GNU's `objdump` tool formats this value as hexadecimal as well, but with spaces, ex. `48 c7 44 24 08 00 00 00 00`.
    * `MOVQ $0x0, 0x8(SP)`
        * The instruction `MOVQ` copies the value from one address to another, ex. `MOVQ SRC, DST`.
        * `MOVQ`
            * Normally `MOVQ` operates right to left, `MOVQ DST SRC`, but as the [Go assembly documentation](https://go.dev/doc/asm) states:

                > One detail evident in the examples from the previous sections is that data in the instructions flows from left to right: MOVQ $0, CX clears CX. This rule applies even on architectures where the conventional notation uses the opposite direction. 
            * The `Q` in `MOVQ` stands for _quadword_:
                * On x86 and x86_64 platforms a _word_ is 16 bits.
                * Because this example is from a 64-bit system, a _quadword_ is 16x4, or...64 bits.
        * `$0x0`
            * The `SRC` of the copy operation.
            * The leading `$` indicates `SRC` is not a memory address, but a literal value.
            * The value to copy is therefore `0x0`, or the integer value `0`.
        * `0x8(SP)`
            * The `DST` of the copy operation.
            * The `0x8` indicates an offset of eight bytes from some address.
            * The address is indicated by `(SP)`, _stack pointer_, the highest memory address of the current stack frame.
            * Therefore `0x8(SP)` can be translated as _eight bytes from the highest memory address of the current strack frame_.

    <br />

    ```
    SP +----------------+ SP + 0 bytes
       |                |
       |                |
       |                |
       |                |
       +----------------+ SP + 8 bytes
       |  name: x       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 0       |
       +----------------+
    ```


1. `xint2yint.go:21	0x2c41			48c7042400000000	MOVQ $0x0, 0(SP)`
    * The assembly for `var y int64`.
    * Copies the literal value `0` to the address `0(SP)`.
        * Please note that `0(SP)` could also be written as `0x0(SP)`.
    * Therefore `0(SP)` can be translated as `_zero bytes from the highest memory address of the current stack frame_.

    <br />

    ```
    SP +----------------+ SP + 0 bytes
       |  name: y       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 0       |
       +----------------+ SP + 8 bytes
       |  name: x       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 0       |
       +----------------+
    ```

4. `xint2yint.go:22	0x2c49			48c744240802000000	MOVQ $0x2, 0x8(SP)`
    * The assembly for `x = 2`
    * `MOVQ $0x2, 0x8(SP)` copies the literal value `2` to the memory address for the variable `x`.

    <br />

    ```
    SP +----------------+ SP + 0 bytes
       |  name: y       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 0       |
       +----------------+ SP + 8 bytes
       |  name: x       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 2       |
       +----------------+
    ```

4. `xint2yint.go:23	0x2c52			48c7042402000000	MOVQ $0x2, 0(SP)`
    * The assembly for `y = x`
    * `MOVQ $0x2, 0(SP)` copies the literal value `2` to the memory address for the variable `y`.
    * Note the compiler, even with optimizations disabled, elected to copy the literal `$0x2` to the address of `y` and not use `MOVQ 0x8(SP) SP` to copy `x` to `y`.

    <br />

    ```
    SP +----------------+ SP + 0 bytes
       |  name: y       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 2       |
       +----------------+ SP + 8 bytes
       |  name: x       |
       |  type: int64   |
       |  size: 8 bytes |
       | value: 2       |
       +----------------+
    ```

In other words, the stack is large enough for two, 64-bit integers. When values are assigned to them, those values are just copied to that variable's address on the stack.
