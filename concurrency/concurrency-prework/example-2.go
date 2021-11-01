package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	done := make(chan struct{}, numTasks)
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Printf("running task %d...\n", i)
			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
