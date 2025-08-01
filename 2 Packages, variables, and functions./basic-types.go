package main

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// rune that is a alias for int32
	s := "¡Hola, 世界!"
	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r)
	}

	// pointer convert to uintptr memory address using unsafe
	var x int = 510
	p := &x                            // puntero a x
	addr := uintptr(unsafe.Pointer(p)) // convertir puntero a entero (dirección de memoria)
	fmt.Println(addr)
}
