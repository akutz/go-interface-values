# Leak

Imagine for a moment there is a kitchen sink with a crack in it. The sink has the _potential_ to leak, but nothing will _escape_ the basin until the sink is used. Much like our imaginary sink, escape analysis inspects variables with the potential to escape when they are used. If that potential exists, the variable is marked as "leaking." 

* [**Criteria**](#criteria): what are the criteria for leaking?
* [**Leak destination**](#leak-destination): where does the water go when it goes down the drain?
  * [**Leaking (to a sink)**](#leaking-to-a-sink): bye-bye
  * [**Leaking to result**](#leaking-to-result): you'll be back!
* [**Leak without escape**](#leak-without-escape): you never left!


## Criteria

There are two requirements to be eligible for leaking:

* The variable must be a function parameter
* The variable must be a reference type, ex. channels, interfaces, maps, pointers, slices

Value types such as built-in numeric types, structs, and arrays are not elgible to be leaked. That does not mean they are never placed on the heap, it just means a parameter of `int32` is not going to send you running for a mop anytime soon.

If the above criteria is met, then a parameter will leak if:

* The variable is returned from the same function and/or
* is assigned to a sink outside of the stack frame to which the variable belongs.


## Leak destination

There are two primary types of leaks:

* leaking (to a sink)
* leaking to result


### Leaking (to a sink)

If a function's parameter is a reference type and the function assigns the parameter to a variable outside of the function, the variable is leaking. While the compiler flag `-m` that prints optimizations does not indicate to _where_ the parameter is leaking, it is helpful to think of this as _leaking to a [sink](https://en.wikipedia.org/wiki/Sink_(computing)). For example:

```golang
var sink *int32

func recordID(id *int32) { // leaking param: id
	sink = id
}
```

The function `recordID` leaks the parameter `id` to the package-level field `sink`.

### Leaking to result

Another type of leak is when a reference parameter is returned from a function:

```golang
func validateID(id *int32) *int32 { // leaking param: id to result ~r1 level=0
	return id
}
```

Other than the fact that `validateID` has very poor validation logic (indeed some might call it non-existent :smiley:), the function returns the `id` parameter. Because the value of `id` is returned, it means it outlives the function's stack frame and has the potential to escape to the heap. Therefore the parameter is marked as _leaking to result_.

## Leak without escape

Please remember, a leak is about the _potential_ to escape to the heap. For example, a parameter can leak to result without ever escaping to the heap. For example:

```go
func main() {
	var id1 int32 = 4096
	if validateID(&id1) == nil {
		os.Exit(1)
	}

	var id2 *int32 = new(int32) // new(int32) does not escape
	*id2 = 4096
	validID := validateID(id2)
	if validID == nil {
		os.Exit(1)
	}
}
```

Use the following command to print the optimizations for the above program:

```bash
docker run -it --rm go-interface-values \
  go tool compile -m ./docs/03-escape-analysis/examples/ex2/main.go
```

The output should resemble:

```bash
./docs/03-escape-analysis/examples/ex2/main.go:22:17: leaking param: id to result ~r0 level=0
./docs/03-escape-analysis/examples/ex2/main.go:32:22: new(int32) does not escape
```

Please note:

* The `id` parameter for the `validateID` function is leaking to result because the function returns the incoming parameter and thus there is potential for the value of `id` to outlive its stack frame.
* The value `id1` is not even mentioned because it is a value type and escape analysis only applies to reference types. While a pointer to `id1` was passed into the `validateID` function, the Go compiler optimized the pointer to the stack.
* The value `id2` _is_ a reference type, but it does not escape. Even though the return value of `validateID` is assigned to `validID`, its object is on the same stack frame as `id2`, thus the latter does not outlive its stack frame. Therefore `id2` does not escape to the heap.


A leak is often only felt a drop at a time, but what about a full-on escape?

---

Next: [Escape](./03-escape.md)
