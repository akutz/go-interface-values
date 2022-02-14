# Escape

Someone turned on the tap and that leaky sink just let water escape. But just like in real life, there does not need to be a pre-existing leak for the floor to get wet. This page discusses when and how variables escape to the heap, either by way of a leaked parameter or all on their own.

* [**Criteria**](#criteria): what are the criteria for escaping to the heap?
* [**Escaping via a leaked parameter**](#escaping-via-a-leaked-parameter): bring the duct tape!
* [**Escaping via a sink**](#escaping-via-a-sink): a little easier to detect


## Criteria

There is one requirement to be eligible for escaping to the heap:

* The variable must be a reference type, ex. channels, interfaces, maps, pointers, slices
* A value type stored in an interface value can also escape to the heap

If the above criteria is met, then a parameter will escape if it outlives its current stack frame. That usually happens when either:

* The variable is sent to a function that assigns the variable to a sink outside the stack frame
* Or the function where the variable is declared assigns it to a sink outside the stack frame


## Escaping via a leaked parameter

A variable can escape via a function with a parameter that leaks to a location outside of the stakc frame. For example:

```go
var sink *int32

//go:noinline
func recordID(id *int32) { // leaking param: id
	sink = id
}

//go:noinline
func main() {
	id1 := new(int32) // new(int32) escapes to the heap
	*id1 = 4096
	recordID(id1)
}
```

Use the following command to print the optimizations for the above program:

```bash
docker run -it --rm go-interface-values \
  go tool compile -m ./docs/03-escape-analysis/examples/ex3/main.go
```

The output should resemble:

```bash
./docs/03-escape-analysis/examples/ex3/main.go:22:15: leaking param: id
./docs/03-escape-analysis/examples/ex3/main.go:28:12: new(int32) escapes to heap
```

Unlike the similar example from the previous page, this time `new(int32)` _does_ escape to the heap. The escape occurs because:

* `main` declared `id1` as an `*int32`
* `id1` is used as an argument for the `recordID` function's `id` parameter
* the `recordID` paramter `id` leaks to the package-level `sink` variable
* package scoped variables are not rooted on the same stack frame as function locals

However, a variable that escapes to the heap does not need to do so via a leaked function parameter.

## Escaping via a sink

You have probably already guessed the other way a variable escapes to the heap -- _What if `main` assigned `id1` to `sink` directly?_ And you know what, you would be right! A variable escapes to the heap if it is a reference type and outlives its stack frame. The purpose of the two, distinct examples on this page is to illustrate that _leak_, _escape_, _move_ is not a sequence -- a variable can escape to the heap without being used with a leaky parameter. Consider the following example:

```go
package main

var sink *int32

//go:noinline
func main() {
	id1 := new(int32) // new(int32) escapes to the heap
	*id1 = 4096
	sink = id1
}
```

Use the following command to print the optimizations for the above program:

```bash
docker run -it --rm go-interface-values \
  go tool compile -m ./docs/03-escape-analysis/examples/ex4/main.go
```

The output should resemble:

```bash
./docs/03-escape-analysis/examples/ex4/main.go:23:12: new(int32) escapes to heap
```

With no leak in sight the value of `id1` still escapes to the heap. I am sure some of you have noticed the strange way the last few examples have assigned the value to the id variables. What happens if we simplify it like so?

```go
id1 := int32(4096)
sink = &id1
```

Well, I suppose you will just have to _move_ to the next page to find out!

---

Next: [Move](./04-move.md)
