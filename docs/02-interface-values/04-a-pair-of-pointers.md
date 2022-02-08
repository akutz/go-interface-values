# A pair of pointers

Through type assertion we have ascertained that storing a value in interface preserves the value's type. That is because under the hood an `interface{}` is really a pair of `uintptr` values:

* an address to the type stored in the interface
* an address to the value stored in the interface

---

:wave: **The last word**

There used to be an optimization where a value stored in an interface was stored directly in the last/second word as long as the size of that value's type was smaller than a `uintptr`. However, this optimization was removed in [github.com/golang/go#8405](https://golang.org/issue/8405) due to the introduction of concurrent garbage collection.

---

There are a couple of ways to access this information:

* The built-in [`println`](https://github.com/golang/go/blob/d588f487703e773ba4a2f0a04f2d4141610bff6b/src/builtin/builtin.go#L261-L266) function
* An [`unsafe.Pointer`](https://pkg.go.dev/unsafe#Pointer) and a `[2]uintptr`
* An [`unsafe.Pointer`](https://pkg.go.dev/unsafe#Pointer) and a `struct{ ptyp, pval uintptr }`

Here is an example that uses all three methods ([Golang playground](https://go.dev/play/p/JRU-xZDNvBf)):

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Store an int64 as an interface value.
	iface := interface{}(int64(2))

	// Get an unsafe pointer for "iface".
	ptrIface := unsafe.Pointer(&iface)

	// Cast the pointer to a *[2]uintptr.
	ptrList := (*[2]uintptr)(ptrIface)

	// Cast the pointer to a *struct{uintptr, uintptr}
	ptrData := (*struct{ ptyp, pval uintptr })(ptrIface)

	// Print the addresses using:
	//   * the builtin println function
	//   * the array of two uintptrs
	//   * the struct with the uintptrs
	println(iface)
	fmt.Printf("(0x%x,0x%x)\n", ptrList[0], ptrList[1])
	fmt.Printf("(0x%x,0x%x)\n", ptrData.ptyp, ptrData.pval)
}
```

The above program produce outputs similar, but not identical, to:

```bash
(0x486920,0x4b1f88)
(0x486920,0x4b1f88)
(0x486920,0x4b1f88)
```

Based on the knowledge of what an interface actually is, we can infer:

* `0x486920` is the address to the type stored in the interface
* `0x4b1f88` is the address to the value stored in the interface

Keep reading to find out how to use the above addresses to access the underlying type and value.

---

Next: [Underlying type](./05-underlying-type.md)
