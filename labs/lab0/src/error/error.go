package main


import (
    "fmt"
    "time"
    "example.com/wmath"
)

// The error type is a built-in interface similar to fmt.Stringer
// type error interface {
//     Error() string
// }

type MyError struct {
    when time.Time
    what string
}

func (e *MyError) Error() string { // the fmt package looks for the error interface
    // Only need to define this Error() method for your type(class...)
    return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
    return &MyError{
	time.Now(),
	"It didn't work",
    }
}

func main() {
    if err := run(); err != nil {
	fmt.Println(err)
    }

    fmt.Println(wmath.WSqrt(2))

    fmt.Println(wmath.WSqrt(-2))
}


