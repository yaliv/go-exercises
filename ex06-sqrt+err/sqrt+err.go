package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("\"cannot Sqrt negative number: %v\"", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 { return 0, ErrNegativeSqrt(x) }
	
	z, z1 := 1.0, 0.0
	for {
		z1 = z - (z*z - x) / (2*z)
		if math.Abs(z-z1) < 1.0e-10 { return z1, nil }
		z = z1
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
