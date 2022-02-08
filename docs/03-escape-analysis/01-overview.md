# Overview

Go tries really, _really_ hard to keep memory on the stack where possible. Memory allocated on the stack relieves pressure on the garbage collector as the memory is cleaned up once the stack on which it is allocated no longer exists.

---

Next: [Leak](./0-leak.md)
