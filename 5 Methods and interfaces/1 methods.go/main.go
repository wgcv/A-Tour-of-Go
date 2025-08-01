package main
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// this is so weird to don't have clases
// v is a Receiver is like self.X in python
func (v Vertex) Abs() float64 {
	// self.X vs v.X
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
