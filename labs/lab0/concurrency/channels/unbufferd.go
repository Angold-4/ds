package main

import "time"
import "fmt"

func main() {
  c := make(chan bool)
  go func() {
    time.Sleep(1 * time.Second)
    <-c // receives from a channel
  } ()
  start := time.Now()
  c <- true // blocks until other goroutine receives 
  // so we actually block the whole main goroutine for one second
  fmt.Printf("send took %v\n", time.Since(start))
}
