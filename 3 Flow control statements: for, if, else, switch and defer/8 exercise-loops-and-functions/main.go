package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	old_z := math.MaxFloat32
	count := 0
	for i := math.MaxFloat32; i > 0.001; i = math.Abs(old_z - z) {
		old_z = z
		z -= (z*z - x) / (2 * z)
		count += 1
	}
	return float64(z), count
}

func main() {
	v := 100.0
	fmt.Println(Sqrt(v))
	fmt.Println(math.Sqrt(v))
}
