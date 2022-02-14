# Move

We have discussed when parameters are leaked and when variables escape to the heap, but what about when the heap is so lovely a variable decides to move there? This page discusses when and why escape analysis moves some variables to the heap.

* [**Criteria**](#criteria): what are the criteria for moving to the heap?
* [**Moving to the heap**](#moving-to-the-heap): when a value on the stack is moved to the heap
* [**Storing value types in interfaces**](#storing-value-types-in-interfaces): when a value type escapes to the heap after all

## Criteria

There are two requirements to be eligible for moving to the heap:

* The variable must be a value type
* A reference to the variable is assigned to a location outside of the local stack frame

What does this look like? It is actually semantically the same as when a variable escapes to the heap -- it all just depends on the initial declaration of that varaible's type...

## Moving to the heap

A value type variable is moved to the heap if a reference to it is assigned to a location outside of the variable's local stack frame. For example:

```go
package main

var sink *int32

//go:noinline
func main() {
	var id int32 = 4096
	sink = &id
}
```

Use the following command to print the optimizations for the above program:

```bash
docker run -it --rm go-interface-values \
  go tool compile -m ./docs/03-escape-analysis/examples/ex5/main.go
```

The output should resemble:

```bash
./docs/03-escape-analysis/examples/ex5/main.go:23:6: moved to heap: id
```

Unlike when `id := new(int32)` escapes to the heap, the `id` in the above example is moved to the heap because it was initially a value type. Thus, if all other conditions are met for escaping and moving to the heap:

* a value _**escapes**_ if it is initially a reference type
* a value _**moves**_ if it is initially a value type

Actually, there is _one_ exception to this rule...


## Storing value types in interfaces

It would be too simple if reference types escaped to the heap and value types moved, wouldn't it? One way a value type is marked as _escaping_ to the heap is when the value is stored in an interface. For example:

```go
package main

var sink interface{}

//go:noinline
func main() {
	var id int32 = 4096
	sink = id
}
```

Use the following command to print the optimizations for the above program:

```bash
docker run -it --rm go-interface-values \
  go tool compile -m ./docs/03-escape-analysis/examples/ex6/main.go
```

The output should resemble:

```bash
./docs/03-escape-analysis/examples/ex6/main.go:24:7: id escapes to heap
```

So let's revise our above rules -- if all other conditions are met for escaping and moving to the heap:

* a value _**escapes**_ if it is initially a reference type or the operation **is** storing a value type in an interface
* a value _**moves**_ if it is initially a value type and the operation is **not** storing a value type in an interface

Okay, so this section has made a _lot_ of assertion using poor analogies about leaky sinks, but do all of these claims...hold water? :smiley:

---

Next: [Tests](./05-tests.md)
