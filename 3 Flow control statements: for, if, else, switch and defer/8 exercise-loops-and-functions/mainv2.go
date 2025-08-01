package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {

	z := 1.0
	count := 0
	var old_z float64
	for {
		z, old_z = z-(z*z-x)/(2*z), z
		if math.Abs(old_z-z) < 0.001 {
			break
		}
		count++
	}
	return z, count
}

func main() {
	v := 10.0
	result, iterations := Sqrt(v)
	fmt.Printf("Sqrt(%v) = %.2f iterations: %v\n", v, result, iterations)
	fmt.Printf("math.Sqrt = %v vs Sqrt: %v\n", math.Sqrt(v), result)

}
