package main

func main() {
  c := make(chan bool)
  go func() {
    // if only this goroutine, after main goroutine
    // has terminated, they will also be terminated
    // (main goroutine won't wating for this goroutine)
    <-c
  } ()

  c <- true // the send is going to block until somebody is ready to recieve it
  <-c // also the receiver is gonna to sleep until it recieve something from that channel
}
