package wmath

import (
    "fmt"
)


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func WSqrt(x float64) (float64, error){
    if x < 0 {
	return 0, ErrNegativeSqrt(x)
    }
    z := float64(1)
    fmt.Printf("Sqrt approximation of %v:\n", z)
    for i := 1; i <= 10; i++ {
	// Newton's method
	z -= (z*z - x) / (2*z)
	fmt.Printf("Iteration %v, value = %v\n", i, z)
    }
    return z, nil
}
