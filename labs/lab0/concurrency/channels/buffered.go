package main

func main() {
  c := make(chan bool, 1)

  go func() {
    // why it won't panic -> do not reach here, then the main goroutine is finished
    println(1)
    k := <-c
    println(k)
  } ()

  // Sends to a buffered channel block only when the buffer is full. 
  // Receives block when the buffer is empty.

  c <- true // send won't be terminated until the buffer size hit its capacity
  v := <-c
  println(2)
  println(v)
}

