package main

import (
	"fmt"
	"math"
)

// non-struct types
type MyFloat float64

// You can extend the methods to non-struct types
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
