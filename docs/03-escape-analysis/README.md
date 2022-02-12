# Escape analysis

Before we dig into why storing a value in an interface can result in a heap allocation, we need to discuss escape analysis:

* **[An overview](./01-overview.md)**: an overview of escape analysis
* **[Leak](./02-leak.md)**: when a value leaks to the heap
* **[Escape](./03-escape.md)**: it doesn't take 600 years, or even 20...
* **[Move](./04-move.md)**: are you too good for your home?
* **[Summary](./05-summary.md)**: what we learned about escape analysis

---

Next: [An overview](./01-overview.md)

typedefs for leaks -- https://github.com/golang/go/blob/master/src/cmd/compile/internal/escape/graph.go

enums for escape  -- https://github.com/golang/go/blob/master/src/cmd/compile/internal/ir/node.go

stack frame -- https://github.com/golang/go/blob/master/src/runtime/stack.go

ast and stack frame life - https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890