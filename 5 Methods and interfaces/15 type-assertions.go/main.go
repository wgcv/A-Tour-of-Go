package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// this is the assertation
	// Extract the value of the interface
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// You check if the interface is of a certain type
	// this doesn't have panic, disco is relax
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// paninc at the disco
	f = i.(float64) // panic
	fmt.Println(f)
}
