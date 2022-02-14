package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func fibonacci() func() int {
    // global variables for specific fibonacci function
    n := 0
    a := 0
    b := 1
    c := a + b

    return func() int { // no name function : func()
	var ret int
	switch { // no argument is just like the if-else stmt
	case n == 0:
	    n++
	    ret = 0
	case n == 1:
	    n++
	    ret = 1
	default:
	    ret = c
	    a = b
	    b = c
	    c = a + b
	}
	return ret
    }
}

func main() {
    hypot := func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
    }

    fmt.Println(hypot(5, 12))
    // functions are values too, they can be passed around just like other values
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))


    f := fibonacci() // a function with closures

    for i := 0; i < 30; i++ {
	fmt.Println(f())
    }


}
