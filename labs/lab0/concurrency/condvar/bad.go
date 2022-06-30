package main

import "sync"
import "time"
import "math/rand"

func main() {
  rand.Seed(time.Now().UnixNano())

  count := 0
  finished := 0
  var mu sync.Mutex

  for i := 0; i < 10; i++ {
    go func() {
      vote := requestVote()
      mu.Lock()
      defer mu.Unlock()

      if vote {
	count++
      }

      finished++
    } ()
  }

  // some threads running concurrently that are making updates to some shared data
  // and I have some other threads that are waiting for some property of 
  // that shared becoming true
  // Clean Sol -> condition variable (see vote-count.go)

  for {
    // problem: busy waiting
    // this for loop will burnning up 100% CPU
    mu.Lock()
    if count >= 5 || finished == 10 {
      break
    }
    mu.Unlock()

    // -> one solution: sleep periodically
    // time.Sleep(50 * time.Millisecond)
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
