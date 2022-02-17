# Lessons learned

Wow, what a ride! What I meant to be a quick detour to better understanding the performance speedup using generics in 1.18 might bring turned into a deep-dive into storing interface values! What did we learn along the way?

* A Go interface is really just a pair of `uintptr` values that point to the interface's underlying type and value.


---

