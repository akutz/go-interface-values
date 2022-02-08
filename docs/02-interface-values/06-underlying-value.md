# Underlying value

Last time on _Go interface values_:

* Go interfaces are really two `uintptr` values that address the underlying type and value stored in the interface
* The address of the type points to an internal Go type where an equivalent to a `reflect.Kind` value is stored
* Type information is global and shared by all values stored in interfaces
* The address of the value may sometimes be duplicated even across distinct types such as `int32` and `int64`?

Okay, now that we are all caught up, let's take a quick look at the example from the previous page ([Golang playground](https://go.dev/play/p/ewZtZafue19)):

```go
package main

func main() {
	println(interface{}(int32(3)))
	println(interface{}(int32(5)))
	println(interface{}(int64(3)))
	println(interface{}(int64(5)))
}
```

We learned the value addresses for `int32(3)` and `int64(3)` are the same as well as `int32(5)` and `int64(5)`:

* **`0x476598`**: value address for `int32(3)` and `int64(3)`
* **`0x4765a0`**: value address for `int32(5)` and `int64(5)`

The reason for this is because we stored _constant_ values in interfaces, and Go's compiler is smart enough to reference the same value multiple times when it is small enough to fit into the width of a larger type (as both `3` and `5` fit into both an `int32` and `int64`). If we were to refactor things a bit so `3` and `5` are no longer constant ([Golang playground](https://go.dev/play/p/qBBwCOhp4II)):

```go
package main

func main() {
	a, b := int32(3), int32(5)
	x, y := int64(3), int64(5)
	println(interface{}(a))
	println(interface{}(b))
	println(interface{}(x))
	println(interface{}(y))
}
```

We see all the value addresses are distinct from one another:

```bash
(0x459d80,0xc00003475c)
(0x459d80,0xc000034758)
(0x459dc0,0xc000034760)
(0x459dc0,0xc000034768)
```

There's _one more thing_. Consider the following example ([Golang playground](https://go.dev/play/p/pAegxrruNvR)):

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

What do you think will happen? Join me on the next page and we will find out together!

---

Next: [Special values](./07-special-values.md)
