package main

import  (
    "fmt"
)

// channels are a typed conduit
// which you can send and recieve values with the channel operator <-

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
	sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)

    // aim: sm the numbers in a slice
    // we distributing the work between two goroutines
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    // once the other side (right) is ready, the value will be sent through channel
    y := <-c
    x := <-c

    fmt.Println(x, y, x+y)
}
