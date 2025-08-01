package main

import (
	"fmt"
	"time"
)

// a normal struct
type MyError struct {
	When time.Time
	What string
}

// with this have the error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// more realistic error is
// func run() (T, error) where T is the result type
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// doesnt exit try catch
	// this is the try catch
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
