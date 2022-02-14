package main

import "golang.org/x/tour/pic"

// https://stackoverflow.com/questions/25459474/tour-of-go-exercise-18-slices
func Pic(dx, dy int) [][]uint8 {
    p := make([][]uint8, dy);
    // p is a slice of slice, with size == dy
    for y := range p {
	// y is a slice index
	p[y] = make([]uint8, dx)
	for x := range p[y] {
	    p[y][x] = uint8(x^y)
	}
    }
    return p
}

func main() {
    pic.Show(Pic)
}
