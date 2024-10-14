package main

import (
	"fmt"
	"sync"
)

func main() {
	limit := 10 // Define the maximum number to print
	oddChan := make(chan int)
	evenChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2) // We have two goroutines

	// Odd number printing goroutine
	go func() {
		defer wg.Done()
		for num := range oddChan {
			if num > limit {
				close(evenChan) // Signal even goroutine to stop
				return
			}
			fmt.Println("odd gr - ",num)    // Print the odd number
			evenChan <- num + 1 // Send the next even number
		}
	}()

	// Even number printing goroutine
	go func() {
		defer wg.Done()
		for num := range evenChan {
			if num > limit {
				close(oddChan) // Signal odd goroutine to stop
				return
			}
			fmt.Println("even gr- ",num)    // Print the even number
			oddChan <- num + 1 // Send the next odd number
		}
	}()

	// Start the sequence by sending the first odd number
	oddChan <- 1

	// Wait for both goroutines to finish
	wg.Wait()
}