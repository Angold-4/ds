package main

import "time"
import "math/rand"

func main() {
  c := make(chan int)

  for i := 0; i < 4; i++ {
    go doWork(c)
  }

  for {
    // The consumer
    v := <-c
    println(v)
  }
}

func doWork(c chan int) {
  // The producer
  for {
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    c <- rand.Int()
  }
}
