# Tests

This section on escape analysis has been rather assertive in its claims about criteria and behaviors. Truthfully this entire repository is the embodiment of "one physical test is worth one thousand expert opinions." I am no expert on, well, much of anything, but I do like digging into problems, even those I do not fully understand. Storing interface values, heap allocation, and escape analysis are such problems. The sources are there in Go, and I can follow _just_ enough to be annoying on Gopher Slack :smiley:

And while I have no _personal_ problem with being incorrect, I am hesitant to assert information as a "source of truth" without a level of certainty I am unable to gleam from the sources. To that end this repository has an obscene number of tests to verify the statements contained within.

* [**Test coverage**](#test-coverage)
* [**List of tests**](#list-of-tests)
* [**Running the tests**](#running-the-tests)
* [**Test framework**](#test-framework)
* [**Writing a test**](#writing-a-test)

## Test coverage

The tests validate the following information:

* The number of expected heap allocations
* The number of expected, allocated bytes
* Which values did not participate in escape analysis, did not escape, leaked, escaped, and were moved

## List of tests

There are number of tests divided into four categories:

* No escape
  * When a pointer does not outlive its call stack
  * A value type that does not escape
  * A value type that does not leak
* Leak
  * To sink
  * A pointer to result
  * A map to result
  * A slice to result
  * A chan to result
  * Five reference values to result with zero mallocs
  * Did not escape because the leaked result did not outlive its stack frame
* Escape
  * 
* Move
  * Pointing to stack from the heap
  * Too large for the stack


## Running the tests

Running the tests is as easy as the following command:

```bash
docker run -it --rm go-interface-values go test -v ./tests/lem
```

## Test framework

The tests are based on a comment-driven framework I assembled called leak, escape, move -- or lem for short. The tests are located in test sources contained in `./test/lem`, for example:

```golang
// lem.move1.name=pointing to stack from heap
// lem.move1.alloc=1
// lem.move1.bytes=4
func move1(b *testing.B) {
	var sink *int32
	for i := 0; i < b.N; i++ {
		var x int32 = 4096 // lem.move1.m=moved to heap: x
		sink = &x
	}
	_ = sink
}
func init() {
	lemFuncs["move1"] = move1
}
```

It should be clear what I meant by _comment-driven framework_. Comments are used to provide the following information:

* [**The test group**](#the-test-group)
* [**The scalar ID of the test within the test group**](#the-scalar-id-of-the-test-within-the-test-group)
* [**The descriptive name of the test**](#the-descriptive-name-of-the-test)
* [**The number of expected allocations**](#the-number-of-expected-allocations)
* [**The number of expected, allocated bytes**](#the-number-of-expected-allocated-bytes)
* [**Zero to many assertions that an optimization should should occur**](#zero-to-many-assertions-that-an-optimizaiton-should-occur)
* [**Zero to many assertions that an optimizaiton should _not_ occur**](#zero-to-many-assertions-that-an-optimizaiton-should-not-occur)

### The test group

The format `<TEST_GROUP><SCALAR_ID>` is used by all comments to uniquely identify the test. The group is also used as the first part of the test's run heirarchy. The descriptive name of the test is the second part of the run heirarchy. 

### The scalar ID of the test within the test group
### The descriptive name of the test
### The number of expected allocations
### The number of expected, allocated bytes
### Zero to many assertions that an optimization should should occur
### Zero to many assertions that an optimizaiton should _not_ occur


## Writing a test

---

Next: [Summary](./06-summary.md)
