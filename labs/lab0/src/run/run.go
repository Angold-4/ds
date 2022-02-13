package main

import (
    "fmt"
    "math"
    "example.com/wmath"
)

func main() {
    guess := wmath.WSqrt(2)
    expected := math.Sqrt(2)

    fmt.Printf("Guess: %v, Expected: %v, Error: %v", guess, expected, math.Abs(guess - expected))
}
