package wmath

import (
    "fmt"
)

func WSqrt(x float64) float64 {
    z := float64(1)
    fmt.Printf("Sqrt approximation of %v:\n", z)
    for i := 1; i <= 10; i++ {
	// Newton's method
	z -= (z*z - x) / (2*z)
	fmt.Printf("Iteration %v, value = %v\n", i, z)
    }
    return z
}
