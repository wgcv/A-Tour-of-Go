package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("mundo")

	fmt.Println("hola")

	fmt.Println("Review Defer  for time")

	// the argument keep the value before the defer is executed but the function run in the end
	defer fmt.Println("Tiempo defer:", time.Now())

	time.Sleep(2 * time.Second)

	fmt.Println("Fin del main:", time.Now())

	// looks like use a LIFO stack

}
