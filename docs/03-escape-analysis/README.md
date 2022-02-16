# Escape analysis

Before we dig into why storing a value in an interface can result in a heap allocation, we need to discuss escape analysis:

* **[An overview](./01-overview.md)**: an overview of escape analysis
* **[Leak](./02-leak.md)**: potential for increased pressure on the garbage collector
* **[Escape](./03-escape.md)**: when go is unable to store a reference type on the stack
* **[Move](./04-move.md)**: when go moves a value type into the heap
* **[Tests](./05-tests.md)**: a test suite to validate what we have learned

---

Next: [An overview](./01-overview.md)
