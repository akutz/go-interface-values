# Type assertion

So far the examples have revolved around storing an `int64` in an interface:

```go
var x int64
var y interface{}
x = 2
y = x
```

We also now know that in order to use `y` in situations where type matters, such as `x + y`, we must assert that `y` _is_ an `int64` just like `x`:

```go
fmt.Println(x + y.(int64))
```

However, the previous page asked the following question: _How do we know the type assertion only worked because the literal value of `2` is a valid value for the type `int64`?_

One way to answer the question is by attempting to assert `y` is a `string` ([Golang playground](https://go.dev/play/p/jEqgcZWXKH9)):

```go
var s string
s = y.(string)
```

The above code compiles fine, but running the example results in a panic:

```bash
panic: interface conversion: interface {} is int64, not string
```

We cannot assert `y` is a `string` because the value stored in `y` is an `int64`. In fact, if you think about it, there is an even stricter constraint if we were to remove interfaces from the equation altogether ([Golang playground](https://go.dev/play/p/w24uFcZppRh)):

```golang
var x int64
var s string
x = 2
s = x.(string)
```

Forget a runtime panic, the above example does not even compile:

```bash
invalid type assertion: x.(string) (non-interface type int64 on left)
```

The above examples demonstrate:

* There are runtime checks to disallow invalid type assertions for interface values.
* There are compile-type checks to disallow invalid type assertions.

Taking these facts into account, we can assume that storing a value in an interface must still preserve the value's type in some way. Keep reading to find out how!

---

Next: [A pair of pointers](./04-a-pair-of-pointers.md)
