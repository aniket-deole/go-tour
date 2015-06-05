package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	i := 0
	z := 1.0
	for ; i < 100; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
