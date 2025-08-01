package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))

	fmt.Println(swap("hello", "world"))
}
