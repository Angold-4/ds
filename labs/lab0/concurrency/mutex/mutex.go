package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
    mu sync.Mutex
    v map[string]int
}

// Inc incrments the counte for the given key
func (c *SafeCounter) Inc(key string) {
     c.mu.Lock()

    // at this time, all goroutines that wants to Lock() this mutex
    // will be put into waiting list
    c.v[key]++;

    c.mu.Unlock();
}

func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()

    // defer means when this program ret, it will execute this stmt
    c.mu.Unlock()

    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
	go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}

