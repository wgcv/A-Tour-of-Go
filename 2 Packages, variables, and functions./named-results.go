package main

import "fmt"

// You can name the return values of a function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// This is a return statement. It returns the values x and y.
	return
}

func main() {
	fmt.Println(split(17))
}
