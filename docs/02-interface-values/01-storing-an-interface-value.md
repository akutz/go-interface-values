# Storing an interface value

The first question people probably have is _What is an interface value?_ To that I say _Dang, I was hoping you would tell me!_. Jokes aside, please consider the following ([Golang playground](https://go.dev/play/p/q6LRdv5H2Rw)):

```go
var x int64
var y interface{}
x = 2
y = x
```

What happens when `y` is assigned the value of `x`? This is known as _storing a value in an interface_, or just _an interface value_.

But how do we _use_ a value once it is stored in an interface? Keep reading to find out!

---

Next: [Using an interface value](./02-using-an-interface-value.md)
