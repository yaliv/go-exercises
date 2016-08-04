package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	/*
	z := 1.0
	for i := 1; i <= 10; i++ {
		z = z - (z*z - x) / (2*z)
	}
	return z
	*/
	
	z, z1 := 1.0, 0.0
	for {
		z1 = z - (z*z - x) / (2*z)
		if math.Abs(z-z1) < 1.0e-10 { return z1 }
		z = z1
	}
}

func main() {
	x := 11111.0
	fmt.Println("math.Sqrt:\t", math.Sqrt(x))
	fmt.Println("Newton's:\t", Sqrt(x))
}
