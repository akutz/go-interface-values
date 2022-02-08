# Special values

Based on everything we have learned so far...

* Type addresses are shared
* Value addresses are shared for constants and distinct for variables

...we _should_ be able to assume the following about the example below ([Golang playground](https://go.dev/play/p/pAegxrruNvR)):

* The type address for `a` and `b` should be the same
* The value addresses for `a`, `b`, `x`, and `y` should all be distinct

```go
package main

func main() {
	a, b := byte(3), byte(5)
	x, y := uint8(3), uint8(5)
	println(interface{}(a))
	println(interface{}(b))
	println(interface{}(x))
	println(interface{}(y))
}
```

Except you have probably realized by now that I would not be posing the question if the answer were so simple :smiley: Here's the output:

```bash
(0x45a240,0x4b89b8)
(0x45a240,0x4b89c8)
(0x45a240,0x4b89b8)
(0x45a240,0x4b89c8)
```

Whoah! Not only is the address to the type the same for `byte` and `uint8`, but the addresses for the values are the same across the types as well! Before I answer the question, I have a magic trick I would like to show you ([Golang playground](https://go.dev/play/p/I-ijkvTlrds)):

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := byte(3)

	// Store x as an interface value.
	iface := interface{}(x)

	// Get an unsafe pointer for "iface".
	ptrIface := unsafe.Pointer(&iface)

	// Cast the pointer to a *[2]uintptr.
	ptrList := (*[2]uintptr)(ptrIface)

	// Get an unsafe.Pointer for the value address offset
	// by the value 197*8.
	ptrY := unsafe.Pointer(ptrList[1] + (197 * 8))

	// Cast ptrY to a uint8.
	y := (*uint8)(ptrY)

	// What is going to be printed?
	fmt.Println(*y)
}
```

So what do you think will be printed? If you guessed `200`, then congratulations, you successfully peeked behind the curtain! It turns out that Go defines a special, internal array with exactly 256 elements called [`staticuint64s`](https://github.com/golang/go/blob/2580d0e08d5e9f979b943758d3c49877fb2324cb/src/runtime/iface.go#L492-L526). This array holds every possible value (0-255) for an integer that is only one byte wide, such as:

* `bool`
* `byte`
* `int8`
* `uint8`

Thus when it is necessary to create a pointer to one of those values, the Go compiler just addresses an element from the aforementioned array. This is also why the above trick was possible:

* By adding `197 * 8` to the address that pointed to the element in the array for `3`
* we addressed the value of `200` from the array
* because each element is eight bytes wide, thus `197 * 8` is the offset of `200` from `3`

Just about now you are probably wondering:

* _If the array is typed `uint64`, does the above work for all integer values?_
* _Why was it necessary to reference the array for the value at all when `x` already has the value?_

The answer to the first question is _no_, to which you will probably ask _Then why is the array `uint64`?_ Great question! We will come back to that later. For now let's keep reading to understand why the value address did not just point to `x`...

---

Next: [Copy on store](./08-copy-on-store.md)
