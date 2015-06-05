package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	// The below line generats an error. Explicit type casting is required.
	// var z int = f
	fmt.Println(x, y, z)
}
