package main

import (
	"fmt"
	"math"
)

// So you can create polymorphic variables
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"} // i = (valor: &T{"Hello"}, tipo: *T)
	describe(i)
	i.M() // llama a (*T).M()

	i = F(math.Pi) // i = (valor: F(math.Pi), tipo: F)
	describe(i)
	i.M() // llama a (F).M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
