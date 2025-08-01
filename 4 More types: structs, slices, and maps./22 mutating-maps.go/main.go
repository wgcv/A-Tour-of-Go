package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// this is so weird
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	// fmt.Println("Answer" in m)
	m["Answer"] = 0

	if value, exist := m["Answer"]; exist {
		fmt.Println("The value:", value)
	} else {
		fmt.Println("The value:", value, "is not present")
	}

}
