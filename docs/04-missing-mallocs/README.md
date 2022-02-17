# Missing mallocs

You are probably saying to yourself "Bury the lead much?" Heh. The thing is, it is really crucial to understand how storing values in an interface and escape analysis works in order to make sense of why _sometimes_ it is possible to store a value in an interface that escapes to the heap...and does not need any new memory to do so.

_Say what!?_ 

Yep, you heard me..._magic_. Okay, not really, but it is due to something about which I certainly had no clue until diving down this rabbit hole. This section reviews:

* [**Shut up and prove it**](./01-shut-up.md): an example of when this situation will occur
* [**Looking at the assembly**](./02-why.md): looking closely at why this happens
* [**When it will happen?**](./03-when.md): all the scenarios where we won't see mallocs
* [**Overall impact**](./04-what.md): what is the impact of this behavior?

---

Next: [Shut up and prove it](./01-shut-up.md)
