package main

import "fmt"


func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
	c <-x
	x, y = y, x+y
    }
    close(c)
}

func main() {
    // the second arg is the size of buffered channel
    ch := make(chan int, 2) 

    ch <- 1
    ch <- 2

    fmt.Println(<-ch)
    fmt.Println(<-ch)

    c := make(chan int, 10)

    fmt.Println(cap(c)) // 10 capacity

    go fibonacci(cap(c), c)

    // if not close explicitly
    // all of the go routines are asleep, a deadlock occurs
    for i := range c {
	fmt.Println(i)
    }
}
