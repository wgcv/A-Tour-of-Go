package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// walkAndClose is a helper function that walks the tree and closes the channel
func walkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go walkAndClose(t1, ch1)
	go walkAndClose(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	// Some test with AI Help
	fmt.Println("Testing Walk function:")
	ch := make(chan int)
	go walkAndClose(tree.New(1), ch)

	fmt.Print("Values from tree.New(1): ")
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("\nTesting Same function:")

	result1 := Same(tree.New(1), tree.New(1))
	fmt.Printf("Same(tree.New(1), tree.New(1)) = %t\n", result1)

	result2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("Same(tree.New(1), tree.New(2)) = %t\n", result2)

	fmt.Println("\nAdditional tests:")
	result3 := Same(tree.New(2), tree.New(2))
	fmt.Printf("Same(tree.New(2), tree.New(2)) = %t\n", result3)

	fmt.Print("Values from tree.New(2): ")
	ch2 := make(chan int)
	go walkAndClose(tree.New(2), ch2)
	for v := range ch2 {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
