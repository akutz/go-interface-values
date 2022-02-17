# Looking at the assembly

The example on this page is based on the source code in [ex1.go](./examples/ex1/ex1.go):

```go
/* line 19 */ var sink interface{}
/* line 20 */ 
/* line 21 */  func ex1() {
/* line 22 */ 	var x int64
/* line 23 */ 	x = 2
/* line 24 */ 	sink = x
/* line 25 */ }
```

With that in mind, let's get started:

1. Build the `ex1` test binary with compiler flags to prevent write barriers (`-wb=false`), inlining (`-l`), and optimization (`-N`). You would never do this in producton, but it makes walking the assembly easier:

    ```bash
    docker run -it --rm -v "$(pwd):/tmp/pkg" go-interface-values \
      go test -gcflags "-wb=false -l -N" \
      -c -o /tmp/pkg/04-missing-mallocs.ex1 \
      ./docs/04-missing-mallocs/examples/ex1
    ```

1. Dump the symbol `ex1$` from the newly built test binary:

    ```bash
    docker run -it --rm -v "$(pwd):/tmp/pkg" go-interface-values \
      go tool objdump -s ex1$ /tmp/pkg/04-missing-mallocs.ex1
    ```

1. The resulting output will be similar (but not identical) to the following:

    ```assembly
    TEXT go-interface-values/docs/04-missing-mallocs/examples/ex1.ex1(SB) /go-interface-values/docs/04-missing-mallocs/examples/ex1/ex1.go
      ex1.go:19		0x4ee440		493b6610		CMPQ 0x10(R14), SP								
      ex1.go:19		0x4ee444		764e			JBE 0x4ee494									
      ex1.go:19		0x4ee446		4883ec20		SUBQ $0x20, SP									
      ex1.go:19		0x4ee44a		48896c2418		MOVQ BP, 0x18(SP)								
      ex1.go:19		0x4ee44f		488d6c2418		LEAQ 0x18(SP), BP								
      ex1.go:20		0x4ee454		48c744240800000000	MOVQ $0x0, 0x8(SP)								
      ex1.go:21		0x4ee45d		48c744240803000000	MOVQ $0x2, 0x8(SP)								
      ex1.go:22		0x4ee466		b803000000		MOVL $0x2, AX									
      ex1.go:22		0x4ee46b		e8d0c4f1ff		CALL runtime.convT64(SB)							
      ex1.go:22		0x4ee470		4889442410		MOVQ AX, 0x10(SP)								
      ex1.go:22		0x4ee475		488d0de4d10000		LEAQ 0xd1e4(IP), CX								
      ex1.go:22		0x4ee47c		48890ddd7c1000		MOVQ CX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink(SB)	
      ex1.go:22		0x4ee483		488905de7c1000		MOVQ AX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink+8(SB)	
      ex1.go:23		0x4ee48a		488b6c2418		MOVQ 0x18(SP), BP								
      ex1.go:23		0x4ee48f		4883c420		ADDQ $0x20, SP									
      ex1.go:23		0x4ee493		c3			RET										
      ex1.go:19		0x4ee494		e8a728f7ff		CALL runtime.morestack_noctxt.abi0(SB)						
      ex1.go:19		0x4ee499		eba5			JMP go-interface-values/docs/04-missing-mallocs/examples/ex1.ex1(SB)
    ```

    Here are the lines on which we want to focus:

    ```assembly
      ex1.go:20		0x4ee454		48c744240800000000	MOVQ $0x0, 0x8(SP)								
      ex1.go:21		0x4ee45d		48c744240803000000	MOVQ $0x2, 0x8(SP)								
      ex1.go:22		0x4ee466		b803000000		MOVL $0x2, AX									
      ex1.go:22		0x4ee46b		e8d0c4f1ff		CALL runtime.convT64(SB)							
      ex1.go:22		0x4ee470		4889442410		MOVQ AX, 0x10(SP)								
      ex1.go:22		0x4ee475		488d0de4d10000		LEAQ 0xd1e4(IP), CX								
      ex1.go:22		0x4ee47c		48890ddd7c1000		MOVQ CX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink(SB)	
      ex1.go:22		0x4ee483		488905de7c1000		MOVQ AX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink+8(SB)
    ```

1. `ex1.go:20		0x4ee454		48c744240800000000	MOVQ $0x0, 0x8(SP)`
    * The assembly for `var x int64`.
    * `MOVQ $0x0, 0(SP)` copies the literal value `0` to the memory address for the variable `x`.

    <br />

    ![Fig.1](https://raw.github.com/akutz/go-interface-values/main/docs/04-missing-mallocs/images/02-why-fig1.svg?sanitize=true)

1. `ex1.go:21		0x4ee45d		48c744240803000000	MOVQ $0x2, 0x8(SP)`
    * The assembly for `x = 2`.
    * `MOVQ $0x2, 0(SP)` copies the literal value `2` to the memory address for the variable `x`.

    <br />

    ![Fig.2](https://raw.github.com/akutz/go-interface-values/main/docs/04-missing-mallocs/images/02-why-fig2.svg?sanitize=true)

1. `ex1.go:22		0x4ee466		b803000000		MOVL $0x2, AX`
    * The assembly for `sink = x`.
    * `MOVL $0x2, AX` copies the literal value `2` into the register `AX`.

1. `ex1.go:22		0x4ee46b		e8d0c4f1ff		CALL runtime.convT64(SB)`
    * Continuing the assembly for `sink = x`.
    * `CALL runtime.convT64(SB)`
        * Ah, this is new!
        * The `CALL` instruction specifies `runtime.convT64(SB)`. The `(SB)` (_static base_) suffix lets us know that `runtime.convT64` is a global symbol and references the memory address for the procedure to call. In order to examine what that symbol is, we can use the following command to print the symbol's assembly:

            ```bash
            docker run -it --rm -v "$(pwd):/tmp/pkg" go-interface-values \
              go tool objdump -s runtime.convT64 /tmp/pkg/04-missing-mallocs.ex1
            ```

            The above command will dump the assembly for the symbol:

            ```assembly
            TEXT runtime.convT64(SB) /usr/local/go/src/runtime/iface.go
              iface.go:378		0x40a940		493b6610		CMPQ 0x10(R14), SP			
              iface.go:378		0x40a944		7657			JBE 0x40a99d				
              iface.go:378		0x40a946		4883ec20		SUBQ $0x20, SP				
              iface.go:378		0x40a94a		48896c2418		MOVQ BP, 0x18(SP)			
              iface.go:378		0x40a94f		488d6c2418		LEAQ 0x18(SP), BP			
              iface.go:379		0x40a954		483d00010000		CMPQ $0x100, AX				
              iface.go:379		0x40a95a		730d			JAE 0x40a969				
              iface.go:380		0x40a95c		488d0d7dac1d00		LEAQ runtime.staticuint64s(SB), CX	
              iface.go:380		0x40a963		488d0cc1		LEAQ 0(CX)(AX*8), CX			
              iface.go:380		0x40a967		eb27			JMP 0x40a990				
              iface.go:383		0x40a969		4889442428		MOVQ AX, 0x28(SP)			
              iface.go:382		0x40a96e		488b1ddbb41e00		MOVQ runtime.uint64Type(SB), BX		
              iface.go:382		0x40a975		b808000000		MOVL $0x8, AX				
              iface.go:382		0x40a97a		31c9			XORL CX, CX				
              iface.go:382		0x40a97c		0f1f4000		NOPL 0(AX)				
              iface.go:382		0x40a980		e81b1c0000		CALL runtime.mallocgc(SB)		
              iface.go:383		0x40a985		488b542428		MOVQ 0x28(SP), DX			
              iface.go:383		0x40a98a		488910			MOVQ DX, 0(AX)				
              iface.go:385		0x40a98d		4889c1			MOVQ AX, CX				
              iface.go:385		0x40a990		4889c8			MOVQ CX, AX				
              iface.go:385		0x40a993		488b6c2418		MOVQ 0x18(SP), BP			
              iface.go:385		0x40a998		4883c420		ADDQ $0x20, SP				
              iface.go:385		0x40a99c		c3			RET					
              iface.go:378		0x40a99d		4889442408		MOVQ AX, 0x8(SP)			
              iface.go:378		0x40a9a2		e899630500		CALL runtime.morestack_noctxt.abi0(SB)	
              iface.go:378		0x40a9a7		488b442408		MOVQ 0x8(SP), AX			
              iface.go:378		0x40a9ac		eb92			JMP runtime.convT64(SB)	
            ```

            So the `CALL` instruction is calling a global function named `convT64` in the Go source code at `runtime/iface.go`.

        * The function [`convT64`](https://github.com/golang/go/blob/c016133c50512e9a83e7442bd7ac614fe7ca62de/src/runtime/iface.go#L378-L386) looks like this:

            ```go
            func convT64(val uint64) (x unsafe.Pointer) {
            	if val < uint64(len(staticuint64s)) {
            		x = unsafe.Pointer(&staticuint64s[val])
            	} else {
            		x = mallocgc(8, uint64Type, false)
            		*(*uint64)(x) = val
            	}
            	return
            }
            ```

        * If we had to summarize this function, we might say: _If `val` is in the array `staticuint64s` then return the address to that element, otherwise copy `val` to the heap and return the address to that location._

        * We will not review all of it, but let's take a look at a few instructions from the assembly for `convT64`:

            * `CMPQ $0x100, AX`
                * The assembly for `if val < uint64(len(staticuint64s)) {`
                * This instruction compares the literal `256` with the value in register `AX`. 
                * The `AX` register is what the compiler used to store the literal `2`, the value of `x`, when we wanted to store `x` in `sink`.
            * `LEAQ runtime.staticuint64s(SB), CX`
                * This instruction loads the address for the global symbol `runtime.staticuint64s` into the register `CX`
            * `LEAQ 0(CX)(AX*8), CX`
                * Stores an address in the register `CX` that is the address already in `CX` offset by the value in the register `AX` times `8`.
                * This is pointer math to get the address offset by 16 bytes from the address of `staticuint64s`.
                * The reason for 16 bytes is because:
                    * the value in `AX` is `2`
                    * which is multiplied by `8`, because `staticuint64s` is an array of `uint64`, which has a size of 8 bytes
                * Since `staticuint64s` is an array of the values 0-255, then the address of `staticuint64s` points to value `0`.
                * Thus `16` bytes offset from `0` will be the second element in the array, which means...
                * The address of the literal `2` is loaded into the register `CX`.
            * _jump to address `0x40a990`_
            * `MOVQ CX, AX`
                * Copies the value in register `CX` (the address of the literal `2` from the global array `staticuint64s`) into register `AX`

1. `ex1.go:22		0x4ee470		4889442410		MOVQ AX, 0x10(SP)`
    * Continuing the assembly for `sink = x`.
    * This one is strange to me, and I'm not exactly sure of its purpose. I plan to follow up with Gopher Slack and see if anyone there knows.
    * The `AX` register contains the addres of the literal `2` from the global array `staticuint64s`.
    * The instruction copies that address at a 16 byte offset from the top of the stack (just after the memory location of variable `x`).

1. `ex1.go:22		0x4ee475		488d0de4d10000		LEAQ 0xd1e4(IP), CX`
    * Continuing the assembly for `sink = x`.
    * `LEAQ 0xd1e4(IP), CX` stores the address of the next CPI instruction in register `CX`.
    * Ultimately what is stored in `CX` is the address of `type.int64`, a global value that specifies the internal type for an `int64`.

1. `ex1.go:22		0x4ee47c		48890ddd7c1000		MOVQ CX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink(SB)`
    * Continuing the assembly for `sink = x`.
    * `MOVQ CX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink(SB)` copies the value in register `CX`, the address of `type.int64`, to the global symbol `sink`, which is the `interface{}` defined in our example.
    * Since `MOVQ` was used we know this is an eight byte wide value.
    * That corresponds to the size of a `uintptr` on a 64-bit platform.
    * Because the destination was the same as the address of `var sink interface{}`, we know the value is assigned as the _type_ address in the interface.

1. `ex1.go:22		0x4ee483		488905de7c1000		MOVQ AX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink+8(SB)`
    * Continuing the assembly for `sink = x`.
    * `MOVQ AX, go-interface-values/docs/04-missing-mallocs/examples/ex1.sink+8(SB)` copies the value in register `AX`, the addres of the literal `2` from the global array `staticuint64s`, to 8 bytes offset from the start of the symbol `sink`, the `interface{}` defined in our example.
    * Because the destination is an 8 byte offset from the address of `var sink interface{}`, we know the value is assigned as the _value_ address in an interface.

---

Next: [When will it happen?](./03-when.md)
