package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"


// Task type
type TaskType int

const (
    Map     TaskType = 1
    Reduce  TaskType = 2
    Done    TaskType = 3
)

// Step 1, implement RPC structs

/*
 * The first RPC
 * idle worker -> coordinator
 * GetTask RPCs are sent from an idle worker to coordinator to ask for the next task to perform
 */

type GetTaskArgs struct{}

type GetTaskReply struct {
    // what type of task is this?
    TaskType TaskType

    // task number of either map or reduce task
    TaskNum int

    // needed for Map (to know which file to write)
    NReduceTasks int

    // needed for Map (to know which file to read)
    MapFile string

    // needed for Reduce (to know how many intermediate map file to read)
    NMapTasks int
}



/*
 * The second RPC
 * finished worker -> coordinator
 * FinishedTask RPCS are sent from an idle worker to coordinator to indicate that a task has been completed
 */

 type FinishedTaskArgs struct {
    // what type of task was the worker assigned?
    TaskType TaskType

    // which task was it ?
    TaskNum int
 }

 // workers don't need to get a reply
 type FinishedTaskReply struct {}


//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.


// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
