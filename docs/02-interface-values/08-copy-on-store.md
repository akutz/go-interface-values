# Copy on store

The examples have gotten pretty complex, but that has been necessary in order to ask this question: _What is being assigned to `y` in the following example?_

```go
var x int64
var y interface{}
x = 2048
y = x
```

Because of how interfaces "wrap" underlying types and values, we might be tempted to say _The variable `y` is assigned `x`._ Except that is not entirely accurate. Let's use a different example:

```go
var x int64
var y int64
x = 2048
y = x
```

Same question. I think we would all answer _The variable `y` is assigned the _**value**_ of the variable `x`_. That would be correct. And we know that modifying `x` later would not affect `y` ([Golang playground](https://go.dev/play/p/jYC8OO5a02a)):

```go
package main

import "fmt"

func main() {
	var x int64
	var y int64
	x = 2048
	y = x
	fmt.Printf("x=%d, y=%d\n", x, y)

	x = 4096
	fmt.Printf("x=%d, y=%d\n", x, y)
}
```

The above example of course emits:

```bash
x=2048, y=2048
x=4096, y=2048
```

So why would think anything different should happen when `y` is an interface ([Golang playground](https://go.dev/play/p/Gcs91tBp09T))?

```go
var y interface{}
```

The program emits the same output as before:

```bash
x=2048, y=2048
x=4096, y=2048
```

This is because storing a value in an interface results in a copy of that value, just like when assigning a value to another variable of the same type. Even if the value is a pointer, you are creating a copy of that pointer.

This _does_ raise the question though -- if an interface is just two addresses to the underlying type and value, where is the copied value of `y`? We know it's:

* not in the interface itself
* not in `staticuint64s` (see [_Special values_](./07-special-values.md)) as the value is greater than `255`

So where is it? It is in one of two locations, and we are about to take a look at the first one...

---

Next: [On the stack](./09-on-the-stack.md)
