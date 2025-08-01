package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = append(s, 100)
	printSlice(s)

	s = append(s, 101)
	printSlice(s)

	s = append(s, 102)
	printSlice(s)
	//  Panic at the disco! You can't do this like in python
	// s = s[2:20]
	// printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
