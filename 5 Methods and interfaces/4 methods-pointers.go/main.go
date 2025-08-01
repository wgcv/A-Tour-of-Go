package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// methods doesn't modify the value of the struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods can modify the value of the struct if you use a pointer
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}
