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

### Runtime Details (paper's Figure 1):
1. **The input and output are stored on a cluster file system called GFS (Google File System), where the part of MapReduce library splites files over many servers, in typically 16MB to 64MB piece. It then starts up many copies of the prgram on a cluster of machines.** 

2. **One of the copies of the program is special -> the master, the rest are workers that are assigned work by the master. Let's say there are M map tasks and R reduce tasks to assign. The master pick idle workers and assigns each one a map task or a reduce task.**

3. **Master gives Map tasks to workers until all `map()`s complete. A worker who is assigned a map task reads the contents of the corresponding input split, and then write output (intermediate data) to local disk.**

4. **Periodically, the intermediate data pairs are written to local disk of each Map workers (in region groups), and the locations of each worker's disk group are passed back to the master, who is responsible for forwarding these locations to the Reduce**

5. **After all Mapes have finished, master hands out Reduce tasks, when a Reduce workers is notified by the master about these locations (different regions), it read the data from disk from map workers. When a Reduce worker has read all intermediate datam it sorts it by the intermediate keys so that all occurrences of the same keys are grouped together**.

6. **The reduce worker iterates over the sorted intermediate data and for each unique intermediate key encounted, it passes the key and the corresponding set of intermediate values to the user's Reduce function, and write the output file on GFS.**

**Finally, when all map tasks and reduce tasks have been compeleted, the master wakes up the user program. At this point, the MapReduce call in the user program returns back to the user code.**

### Fault Tolerance

Since the MapReduce library is designed to help process very large amounts of data using hundreds or thousands of machines, the library must tolerate machine failures gracefullly. (Must insert fault tolerance into its design)


The master keeps several data structures.<br>
**For each map task and reduce task. it stores the state (idle, in-progress, or completed), and the identity of the worker machine.**

