package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a_o, b_o := "hello", "world"
	fmt.Println(a_o, b_o)
	a, b := swap(a_o, b_o)
	fmt.Println(a, b)

}
