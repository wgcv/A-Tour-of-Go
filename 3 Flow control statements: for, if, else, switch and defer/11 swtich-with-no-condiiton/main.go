package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// is like a if else doesn't review the switch only each case
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
