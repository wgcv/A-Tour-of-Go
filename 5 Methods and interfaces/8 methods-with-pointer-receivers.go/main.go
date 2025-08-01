package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// A good practice is to use only pointer receivers or only value receivers, but not both.
// For big structs, it's better to use pointer receivers to avoid copying the whole struct.
// For small structs, it's better to use value receivers to avoid the overhead of allocating memory for the pointer.
// a small struct is a struct that has less than 10 basic types (int, float, string, etc.) -- defined by my
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
