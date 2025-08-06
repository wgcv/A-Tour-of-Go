package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// if you don't wait for the goroutine to finish,
		// the main function will finish before the goroutine and
		// the output will be "hello" only.
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	go say("world")
	say("hello")
}
