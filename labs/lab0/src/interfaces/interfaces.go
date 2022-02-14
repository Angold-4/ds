package main

import (
    "fmt"
    "math"
)

// An interfaces type is defined as a set of method signatures
// A type implements an interace by implementing its methods 
// (a type is just like a class with properties (its variables))

type I interface {
    M()
}

type T struct {
    // just like properties (struct)
    S string
}

type F float64


// This methed means type T implements the interface I
func (t *T) M() {
    // M has a pointer reciever
    fmt.Println(t.S)
}

// This method means type F also implements the interface I
func (f F) M() {
    // M has a literal value reciever (type)
    fmt.Println(f)
}


func main() {
    var i I = &T{"hello"} // since the method M has a pointer receiver
    // we can't define T(literal value) as an I value
    i.M() // hello
    describe(i) // (&{hello}, *main.T) 

    i = F(math.Pi)
    describe(i) // (3.141592653589793, main.F)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
