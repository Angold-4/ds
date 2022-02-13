package main

import (
    "fmt"
)

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
    // name is a array
    names := [4]string {
	"John",
	"Paul",
	"George",
	"Ringo",
    }

    fmt.Println(names)

    a := names[0:2] // a slice
    b := names[1:3] // a slice

    fmt.Println(a, b)

    // slice are like reference to arrays
    // a slice does not store any data, 
    // it just describes a section of an underlying array

    b[0] = "Angold"
    fmt.Println(a, b)  // [John Angold] [Angold George]
    fmt.Println(names) // [John Angold George Ringo]

    // A slice has both a length and capacity
    // length:  # of elements it contains
    // capaciy: # of elements in the underlying array
    // counting from the first element in the slice

    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    s = s[:0] // exclued 0, so the length becomes 0
    printSlice(s);

    s = s[:4]
    printSlice(s);

    s = s[2:6]
    printSlice(s);
}

