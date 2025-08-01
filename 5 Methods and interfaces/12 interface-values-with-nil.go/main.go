package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
	var t2 *T = nil // t2 es un puntero nil de tipo *T
	var i2 I = t2   // i2 guarda (valor=nil, tipo=*T) → i2 NO es nil

	var j2 I // j2 es nil, porque no apunta a nada ni tiene tipo
	describe(i2)
	describe(j2)
	j2.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
