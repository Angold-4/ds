package main

import "sync"
import "time"
import "math/rand"

func main() {
  rand.Seed(time.Now().UnixNano())

  count := 0
  finished := 0
  var mu sync.Mutex
  cond := sync.NewCond(&mu) // Condvar

  for i := 0; i < 10; i++ {
    go func() {
      vote := requestVote()
      mu.Lock()
      defer mu.Unlock()
      if vote {
	count++
      }
      finished++
      // affect the condition
      cond.Broadcast() // wakes up whoever waiting on the condvar
    } ()
  }

  mu.Lock()
  for count < 5 && finished != 10 {
    cond.Wait()
  }

  if count >= 5 {
    println("received 5+ votes!")
  } else {
    println("lost")
  }
  mu.Unlock()
}

func requestVote() bool {
  // random vote
  time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
  return rand.Int() % 2 == 0
}
