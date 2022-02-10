# 2. RPC and Threads

##### 09/02/2022 By Angold Wang

## 1. RPC and Threads in Golang

### 1. Why Go?
* Good support for threads (go routine)
* Convenient RPC
* Threads + GC is particularly attractive
* Relatively Simple (easy to debug)

### 2. Threads

##### "Allow different parts of the program to sort of be in its own point in a different activity."
##### Each thread includes some per-thread state: `(pc, reg, stack)`.

* [x] **I/O concurrency**
    * While some threads are waiting for I/O, other threads can utilize the CPU resources.

* [x] **Parallelism**
    * Execute code in parallel on several cores.

* [x] **Convinience**
    * relatively simple to programming (compare with event-driven process)


To have a better understanding of threads, let's compare to its alternative:

#### Event-driven Process

##### "Way that write code that explicitly interleaves activities, in a single thread."

* **Implementation Details:**
    * keep a table of each of state about each activity
    * One "event" loop that: 
        1. Checks for new input for each activity
        2. Does the next step for each activity
        3. Update the state table
* [x] **I/O concurrency**
    * The program can pick current unblocked activity to run (just like in threads)
    * It can eliminates thread costs (which can be substantial)
* [ ] **Parallelism**
    * There are no parallelism in multi-core machine (Only one thread)

* [ ] **Convenience**
    * It is usually painful to program, compare to implement with thread

#### Multiple Process `(folk)`

##### "Folk multiple process, just runs like multiple threads, which will gives you thread-like I/O concurrency, and parallelism."

Although in user space, thread seems look like is implemented by the language itself, i.e. Some OS books says "the kernel is not aware of the existence of threads."<br>
**In my opinion, thread is some sort of the minimum-executable unit in the machine.<br>** 
**Which means, in some low-level perspective, a thread has no difference than a process, They are all just executable routines with its own states.**<br>

> i.e. For instance, think of a multi-core machine, when there are multiple threads runs on it, it just like multiple processes with same address space.

##### From an Operating System point of view, there are two stages of scheduling: First, the Process scheduler pick a process to run, which contains one or more routines and its own address space. Second, for each process, the written language implements the details of their execution (threads or whatever).

Compare to multi-threading, when there are multiple processes runs in different cores, the communication between them costs much.(compare to threads, since threads are on the same address space.) And the context switching between each process also costs much (OS sheduler v.s. Thread scheduler in language).


##### Threading Challenges

Threads are convenient, because a lot of times they allow you to write the code for each thread just as if it were a pretty ordinary sequenial program.<br>

Since all threads in the same process are in the same address space. There are also some unexpected errors occurs:

* Race Condition
* Coordination
* Dead Lock















