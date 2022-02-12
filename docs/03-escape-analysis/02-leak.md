# Leak

Escape analysis marks a variable as "leakd" when there is _potential_ for an object to outlive a stack frame:

* Value type parameters cannot be marked as leaking:
    * `int`
* Non-value type parameters can be marked as leaking:
    * `Map`
    * `Pointer`
    * `Slice`


For instance, let's look at the `Login` function from [`./examples/lem/main.go`](../../examples/lem/main.go):

```go
/* 19 */ var lastLoginAttemptUsername *string
/* 20 */ 
/* 21 */ func Login(
/* 22 */     username *string,
/* 23 */     password *string,
/* 24 */     token interface{},
/* 25 */ ) interface{} {
/* 26 */ 
/* 27 */     lastLoginAttemptUsername = username
/* 28 */     switch {
/* 29 */     case username != nil && password != nil:
/* 30 */         return "newLoginToken"
/* 31 */     case token != nil:
/* 32 */         return token
/* 33 */     default:
/* 34 */         return nil
/* 35 */     }
/* 36 */ }
```

We can use the `go tool compile` command to see what has potential for heap allocation:

```bash
docker run -it --rm go-interface-values go tool compile -l -N -m ./examples/lem/main.go
```

The revelant lines from the output are:

```bash
./examples/lem/main.go:22:2: leaking param: username
...
./examples/lem/main.go:24:2: leaking param: token to result ~r0 level=0
```



1. Why is `username` leaking?
2. Why is `token` leaking?
3. What does `to result` mean, and why does `token` have it and `username` does not?

The parameter

---

Next: [Escape](./03-escape.md)
