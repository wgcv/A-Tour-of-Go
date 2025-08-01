package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		// go can manage complex numbers
		return sqrt(-x) + "i"
	}
	// convert int to string
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
