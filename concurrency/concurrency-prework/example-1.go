package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		currentVal := i // Without pulling out this value i keeps incrementing and all print statements say 10 or close to it
		go func() {
			fmt.Printf("launched goroutine %d\n", currentVal)
		}()
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
