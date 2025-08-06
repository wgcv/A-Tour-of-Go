package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Create a channel ever n time
	tick := time.Tick(100 * time.Millisecond)
	// You also can use time.NewTicker and .Stop() prevent memory leak if you stop using

	// The same create a channel but only run one
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
