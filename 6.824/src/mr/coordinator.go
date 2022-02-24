package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"
import "sync"
import "time"



type Coordinator struct {
	// Your definitions here.

	// protect coordinator state 
	// from concurrent access
	mu sync.Mutex


	// Allow coordinator to wait to assign reduce tasks untl map tasks have finished
	// or when all tasks are assigned and are running
	// The coordinator is woken up either when a task has finished, or if a timeout has expired.
	cond *sync.Cond


	mapFiles       []string
	nMapTasks      int
	nReduceTasks   int

	// Keep track of when tasks are assigned
	// and which tasks have finished
	mapTasksFinished      []bool
	mapTasksIssued        []time.Time
	reduceTasksFinished   []bool
	reduceTasksIssued     []time.Time

	// set to true when all reduce tasks are complete
	isDone bool
}


// Your code here -- RPC handlers for the worker to call.


// Handle GetTask RPCs from worker

func (c *Coordinator) HandleGetTask(args *GetTaskArgs, reply *GetTaskReply) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    reply.NMapTasks = c.nMapTasks;
    reply.NReduceTasks = c.nReduceTasks;

    // issue map tasks until there are no map tasks left
    for {
	mapDone := true
	for m, done := range c.mapTasksFinished {
	    // check all map tasks
	    if !done {
		// Assign a task if it is either never been issued, or if it's been
		// too long since it was issued so the worker may have crashed.
		// Note: if task has never been issued, time is initialized to 0 UTC
		if c.mapTasksIssued[m].IsZero() ||
		    time.Since(c.mapTasksIssued[m]).Seconds() > 10 {
			// timeout or not assigned
			reply.TaskType = Map
			reply.TaskNum = m
			reply.MapFile = c.mapFiles[m]
			c.mapTasksIssued[m] = time.Now()
			return nil
		} else {
		    // this map task is in progress
		    mapDone = false
		}
	    }
	}

	// if all maps are in progress and haven't timed out, wait to give another task
	if !mapDone {
	    c.cond.Wait()
	} else {
	    // we're done all the map tasks
	    break
	}
    }

    // All map tasks are done, issue reduce task now
    for {
	redDone := true
	for r, done := range c.reduceTasksFinished {
	    if !done {
		// Assign a task if it is eitehr never been issued, or if it is been too long
		// since it was issued so the worker may have crashed
		// just like the map above
		if c.reduceTasksIssued[r].IsZero() ||
		    time.Since(c.reduceTasksIssued[r]).Seconds() > 10 {
			reply.TaskType = Reduce
			reply.TaskNum = r
			c.reduceTasksIssued[r] = time.Now()
			return nil
		} else {
		    redDone = false
		}
	    }
	}
	// if all reduces are in progress and haven't timed out
	if !redDone {
	    c.cond.Wait()
	} else {
	    // we are done with all reduce tasks
	    break
	}
    }

    reply.TaskType = Done
    c.isDone = true
    return nil
}


func (c *Coordinator) HandleFinishedTask(args *FinishedTaskArgs, reply *FinishedTaskReply) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    switch args.TaskType {
    case Map:
	c.mapTasksFinished[args.TaskNum] = true
    case Reduce:
	c.reduceTasksFinished[args.TaskNum] = true
    default:
	log.Fatalf("Bad finished task? %v", args.TaskType)
    }

    // wake up the GetTask handler loop: a task has finished, so we might be able to assign another one
    c.cond.Broadcast()

    return nil
}


// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}


//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {

	// Register publishes the receiver's methods in the DefaultServer.
	rpc.Register(c)

	rpc.HandleHTTP()

	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	// Listen specify the addr
	l, e := net.Listen("unix", sockname) // l stands for listener 
	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil) // k
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.isDone
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.
	// init
	c.cond = sync.NewCond(&c.mu)
	c.mapFiles = files
	c.nMapTasks = len(files)
	c.mapTasksFinished = make([]bool, len(files))
	c.mapTasksIssued = make([]time.Time, len(files)) // per file a task

	c.nReduceTasks = nReduce
	c.reduceTasksFinished = make([]bool, nReduce)
	c.reduceTasksIssued = make([]time.Time, nReduce)

	// wake up the GetTask handler thread every once in a while to check if some task hasn't
	// finished, sos we can know to reissue it 
	// every second, check whether finished
	go func() {
	    for {
		c.mu.Lock()
		c.cond.Broadcast()
		c.mu.Unlock()
		time.Sleep(time.Second)
	    }
	}()

	c.server()
	return &c
}
