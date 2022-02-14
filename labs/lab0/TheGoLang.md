# The Go Programming Language

![fiveyears](Sources/fiveyears.jpeg)

##### Notes By Angold Wang for the Google Tech Talks on [October 30 2009](https://www.youtube.com/watch?v=rKnDgT73v8s)
##### Presented by Rob Pike.

## 1. Overview

**[Robert Griesmer](https://en.wikipedia.org/wiki/Robert_Griesemer), [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson) and [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike) started the project in late 2007.**<br>

> **"You can be productive or you can be safe, but you can't really be both."**

**The goal of golang: They want the efficiency of a statically compiled language, but the ease of programming of dynamic language.**


* Fundamentals:
    * **Clean, concise syntax.**
    * **Lightweight type system.**
    * **No implicit conversions: keep things explicit.**
    * **Untyped unsized constants. (i.e. no more `0x80ULL`)**
    * **Strict separation of interface and implementation.**

* Runtime:
    * **Garbage collection.**
    * **Strings, maps, communication channels.**
    * **Concurrency** (GC makes concurrency works very well)

* Package model:
    * **Explicit dependencies to enable faster builds.**







