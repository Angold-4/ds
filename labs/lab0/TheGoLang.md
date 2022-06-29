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


## 2. Quotes

* **Way I like to think of this entire world of programming is a vector space of very high dimension, and what you want to do is define the basis set that covers that vertor space, so that you can write the program you want by combining the apporpriate orthogonal set of features.** And what happends when you add features for expressiveness or for fun is tends to add more non basis vectors into that space and so there become many paths to get to a particular solution. [by Rob Pike]

* As a programmer, we tend to spend way more time reading other people's code than writing code.


