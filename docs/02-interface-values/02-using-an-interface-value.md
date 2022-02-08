# Using an interface value

The previous page showed how to store a value in an interface, but how do we use the value stored in `y`? There are _some_ operations that we can perform without needing to know what the value of `y` is or its type, ex. ([Golang playground](https://go.dev/play/p/TULVVw36Q82)):

```go
var x int64
var y interface{}
x = 2
y = x
fmt.Println(x, y)
```

The above code prints the expected output:

```bash
2 2
```

But what about an operation where the type of multiple operands matter, such as the sum of `x` and `y` ([Golang playground](https://go.dev/play/p/EOSUArFfjIP))?

```go
fmt.Println(x + y)
```

Instead of emitting `4`, the program will fail to compile with the following error:

```bash
invalid operation: x + y (mismatched types int64 and interface {})
```

In order to sum `x` and `y`, they must be the same type. Let's try the example one more time ([Golang playground](https://go.dev/play/p/izfHznJns2F)):

```go
fmt.Println(x + y.(int64))
```

Now the code compiles, runs, and emits the expected output of `4`.

The syntax `y.(int64)` is known as _type assertion_. Yet, how do we know the type assertion only worked because the literal value of `2` is a valid value for the type `int64`?

Please continue to the next page to learn how type assertion reveals the inner-workings of interface values...

---

Next: [Type assertion](./03-type-assertion.md)
