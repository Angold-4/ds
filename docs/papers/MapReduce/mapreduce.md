# MapReduce Notes

**By [Angold Wang](https://github.com/Angold-4) | [Original Paper](mapreduce.pdf)**


## 1. Introduction
Back to 2004, engineers in Google have implemented hundreds of special-purpose computations that process large amounts of raw data. Such as building a index of the web, and handle crawled documents, web requests logs, etc.<br>
**However, the input data is usually large and the computations have to be distributed across hundreds or thousands of machines in order to finish in a reasonable amount of time.**<br> 
**How to parallelize the computation, distribute the data, and handle failures**, soon becames a big question.<br>

So they really needed some **kind of framework** that would make it easy to just have their engineers write the kind of guts of whatever analysis they wanted to do, like sort algorithm or a web index or link analyzer or whatever. <br>
**Just write the guts of that application and be able to run it on a thousands of computers without worrying about the details of how to spread the work over the thousands of computers how to organize whatever data movement was required, or how to cope with the inevitable failures.**

#### In short, they were looking for a framework that would make it easy for non-specialists to be able to write and run giant distributed computations.

* Multi-hour Computations
* Multi-terabyte Datasets

## 2. Programming Model

**In a programmer's view, the computation recieves many files, and takes a set of input key/value pairs, then produces a set of output key/value pairs.**

### Example: word count
![mapreduceex](../../Sources/mapreduceex.png)

#### Abstract view of a MapReduce job:

#### Map
* Written by the User.
* Recieve input files and then produces a set of output key/value pairs.
* MR calls `map(k, v)` for each input file `k`, and its contains `v`, produces set of `k2`, `v2` " intermediate" data.
```javascript
map(String key, String value) {
    // key: document name
    // value: document contents
    for each word w in value:
        Emit(w, "1");
}
```


#### Reduce
* Wirtten by the user.
* Recieve key/value pairs from **`Map`**.
* Then merge together these values with same key to **form a possibly smaller set of values**.
* MR gathers all intermediate `v2` for a given `k2`, and passes each key + values to a reduce call.
* Final outut is set of <k2, v3> pairs from `reduce(k2, v2)`s.
```javascript
reduce(k, v)
    // key: a word
    // value: a list of counts
    int result = 0; 
    for each v in values:
        result += ParseInt(v);
    Emit(AsString(result));
```
```javascript
// or in this way:
reduce(k, v)
    emit(len(v));
```

#### MapReduce scales well
* N "worker" computers get you Nx throughtput.
* `map`s can run in parallel, since they don't interact. Same for `reduce`s.
* In that way, you can get more throughtput by buying more computers.


## 3. Implementation Details

![mapreduce](../../Sources/mapreduce.png)

### 1. Runtime Details (paper's Figure 1):
1. **The input and output are stored on a cluster file system called GFS (Google File System), where the part of MapReduce library splites files over many servers, in typically 16MB to 64MB piece. It then starts up many copies of the prgram on a cluster of machines.** 

2. **One of the copies of the program is special -> the master, the rest are workers that are assigned work by the master. Let's say there are M map tasks and R reduce tasks to assign. The master pick idle workers and assigns each one a map task or a reduce task.**

3. **Master gives Map tasks to workers until all `map()`s complete. A worker who is assigned a map task reads the contents of the corresponding input split, and then write output (intermediate data) to local disk.**

4. **Periodically, the intermediate data pairs are written to local disk of each Map workers (in region groups), and the locations of each worker's disk group are passed back to the master, who is responsible for forwarding these locations to the Reduce**

5. **After all Mapes have finished, master hands out Reduce tasks, when a Reduce workers is notified by the master about these locations (different regions), it read the data from disk from map workers. When a Reduce worker has read all intermediate datam it sorts it by the intermediate keys so that all occurrences of the same keys are grouped together**.

6. **The reduce worker iterates over the sorted intermediate data and for each unique intermediate key encounted, it passes the key and the corresponding set of intermediate values to the user's Reduce function, and write the output file on GFS.**

**Finally, when all map tasks and reduce tasks have been compeleted, the master wakes up the user program. At this point, the MapReduce call in the user program returns back to the user code.**

### 2. Fault Tolerance

Since the MapReduce library is designed to help process very large amounts of data using hundreds or thousands of machines, the library must tolerate machine failures gracefullly. (Must insert fault tolerance into its design)


The master keeps several data structures.<br>
**For each map task and reduce task. it stores the state (idle, in-progress, or completed), and the identity of the worker machine.**

#### Worker Failure

##### We want to completely hide failures from the application programmer.

**The master pings every worker periodically. If no response is recieived from a worker in a certain amount of time, the master marks the worker as failed. MR re-runs just the failed `map`s and `reduce`s.**

#### Details of worker crash recovery

#### Map worker crashes:
* **Master notices worker no longer responds to pings.**
* **Master knows which map tasks it ran on the worker, since their output is stored on the failed machine and is there inaccessible by other Reduce workers (no write location back to the master), and the output must be re-executed.**
* **Then Master schedule another worker to run those failed tasks**
* **If some `reduce`s already fetched the intermediate data (if this `map` is already finished), the `reduce` need to be re-running**

#### Reduce worker crashes:
* **Master re-starts worker's unfinished tasks on other workers.**


#### Master Failure
* **During the runtime of Master, the master program will write periodic checkpoints to the master data structures.**
* **If the master task dies, a new copy can be started fromm the last checkpointed state.**


### 3. Locality

##### What will likely limit the performance of MapReduce?
**In 2004, authors were limited by network capacity, the paper's root switch: 100 to 200 gigabits/second, total 1800 machines, so 55 megabits/second/machine. which is relatively small (e.g. much less than disk or RAM speed)**

##### How MR minimize network use?
* **Master tries to run each Map task on GFS Server where stores its input. (i.e. all computers run both GFS and MR workers) so input is read from local disk, not over network.**
* **Intermediate data goes over network just once. which from Map workers' local disk -> Reduce workers, rather than using GFS.**


### 4. Task Granularity

##### How does MR get good load balance?

**Wasteful and slow if N-1 servers have to wait for 1 slow server to finish. But the fact is that some tasks likely take longer than others.**<br>

**We subdivide the map phase into M pieces and the reduce phase into R pieces. Ideally, M and R should be much larger than the number of worker machines. (i.e. many more tasks than workers).<br>**

**Master hands out new tasks to workers who finish previous tasks, so "faster" servers do more than "slower" ones, finish about the same time(hopefully), which improves dynamic load balancing.**<br>

But there are also some limitations of how large M and R can be, some kind of "trade-off", since the Master must make `O(M + R)` scheduling decisions.<br>
In practice, we often perform MapReduce computations with M = 200,000 and R = 5,000 using 2,000 worker machines.


### 5. Backup Tasks
**One of the common causes that lengthens the total time taken for a MapReduce operation is a "straggler": a machine that takses an unusually long time to complete its works.**<br>

**Which means, a machine can compelete its works, but costs too much times than other machine doing the same job. The reason of that is usually some correctable errors that slow its performance, like disk problem, or the competition for CPU, memory, network bandwidth.**<br>

**The engineers in Google have a general mechanism to alleviate the problem of stragglers: When a MapReduce operation is close to completion, the Master schedules backup executions of the remaining** *in-progress* **tasks. The task is marked as completed whenever either the primary of the backup execution completions.**





