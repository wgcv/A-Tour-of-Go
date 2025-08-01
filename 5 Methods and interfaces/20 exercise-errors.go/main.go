package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	old_z := math.MaxFloat32
	count := 0
	for i := math.MaxFloat32; i > 0.001; i = math.Abs(old_z - z) {
		old_z = z
		z -= (z*z - x) / (2 * z)
		count += 1
	}
	return float64(z), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
