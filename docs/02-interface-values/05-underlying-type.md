# Underlying type

Now that we know interfaces are just a pair of pointers with addresses to the underlying type and value, it should be pretty simple to access the values at those addresses, right? Well, it's simple to access the _memory_, but to access the value we need to know the type of the value stored _in_ that memory. Remember, an interface is a tuple with:

* an address to the type stored in the interface
* an address to the value stored in the interface

Therefore if we can figure out how to make sense of the first element, the underlying type, we are a step closer to making sense of the second, the underlying value. Luckily it's fairly straight-forward to do so. An interface stores type information in an internal type, [`type _type struct`](https://github.com/golang/go/blob/d588f487703e773ba4a2f0a04f2d4141610bff6b/src/runtime/type.go#L30-L51) from the [`runtime`](https://pkg.go.dev/runtime) package:

```go
type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      tflag // uint8
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool // uintptr
	gcdata     *byte
	str        nameOff // int32
	ptrToThis  typeOff // int32
}
```

The value of the `kind` field is the one we need, and in fact this value maps directly to the public [`reflect.Kind`](https://pkg.go.dev/reflect#Kind) type. There are two ways to obtain the information we need:

* A modified version of the private `_type` struct
* A memory offset

Here is an example that uses both methods ([Golang playground](https://go.dev/play/p/ddUPcDsDaJQ)):

```go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// _type is a copy of the private runtime._type struct
// (https://bit.ly/34z6d0R) with the following changes:
//
//   * the field "tflag" is changed to its underlying type, uint8
//   * fields after "kind" are omitted
type _type struct {
	size       uintptr // 8 bytes on a 64-bit platform
	ptrdata    uintptr // 8 bytes on a 64-bit platform
	hash       uint32  // 4 bytes
	tflag      uint8   // 1 byte
	align      uint8   // 1 byte
	fieldAlign uint8   // 1 byte
	kind       uint8   // offset by 23 bytes on 64-bit platforms
}

func main() {
	// Store an int64 as an interface value.
	iface := interface{}(int64(2))

	// Get an unsafe pointer for "iface".
	ptrIface := unsafe.Pointer(&iface)

	// Cast the unsafe pointer to a *[2]uintptr.
	ptrList := (*[2]uintptr)(ptrIface)

	// Read the type information using:
	//   * the _type struct
	//   * the memory offset for the kind field -- 23 bytes on 64-bit platforms
	kindFromStruct := reflect.Kind(((*_type)(unsafe.Pointer(ptrList[0]))).kind)
	kindFromOffset := reflect.Kind(*(*uint8)(unsafe.Pointer(ptrList[0] + 23)))

	fmt.Println(kindFromStruct)
	fmt.Println(kindFromOffset)
}
```

The output from the above program will be:

```bash
int64
int64
```

In fact, there is something else interesting about the underlying types -- they are shared across all interfaces. Check out this example ([Golang playground](https://go.dev/play/p/ewZtZafue19)):

```go
package main

func main() {
	println(interface{}(int32(3)))
	println(interface{}(int32(5)))
	println(interface{}(int64(3)))
	println(interface{}(int64(5)))
}
```

The output this time will resemble:

```bash
(0x459d80,0x476598)
(0x459d80,0x4765a0)
(0x459dc0,0x476598)
(0x459dc0,0x4765a0)
```

Huh, it looks like there is some duplicate information. Notice how the addresses for the types are the same for the two `int32` values and the two `int64` values? 

* **`0x459d80`**: type address for `int32(3)` and `int32(5)`
* **`0x459dc0`**: type address for `int64(3)` and `int64(5)`

This is because type information is global and is shared for all values stored in interfaces. However, are the addresses for the values _also_ shared -- across _type_!?

* **`0x476598`**: value address for `int32(3)` and `int64(3)`
* **`0x4765a0`**: value address for `int32(5)` and `int64(5)`

How can this be? Keep reading to find out!

---

Next: [Underlying value](./06-underlying-value.md)
