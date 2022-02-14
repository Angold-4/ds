package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// Go does not have classes. 
// But you can define methods on types
// A method is a fnction with a special receiver argument
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) { // pointer receivers
    // which allow methods modify their receivers
    // pointer receivers are more efficient (like ref or pointer in c++)
    // (only when the receiver is a larget struct)
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    p := &v
    fmt.Println(v.Abs()) // 50

    p.Scale(10)          // Will be interpreted as (*p).Scale()
    fmt.Println(v.Abs()) // 500
}
