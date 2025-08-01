package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	var i int = 42
	var d float64 = float64(i)
	var u uint = uint(d)
	fmt.Println(i, d, u)
}
