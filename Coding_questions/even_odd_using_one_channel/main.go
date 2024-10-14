package main

import (
	"fmt"
	"sync"
)

func main() {
	limit := 10
	var wg sync.WaitGroup
	wg.Add(2) // Two goroutines: one for odd, one for even

	// Control variable to indicate whose turn it is
	type Turn struct {
		num  int
		turn string // "odd" or "even"
	}
	turnChan := make(chan Turn)

	// Odd number printing goroutine
	go func() {
		defer wg.Done()
		for t := range turnChan {
			if t.num > limit {
				close(turnChan)
				return
			}
			if t.turn == "odd" && t.num%2 != 0 {
				fmt.Println(t.num)
				turnChan <- Turn{num: t.num + 1, turn: "even"}
			} else {
				// Pass it to the other goroutine
				turnChan <- t
			}
		}
	}()

	// Even number printing goroutine
	go func() {
		defer wg.Done()
		for t := range turnChan {
			if t.num > limit {
				close(turnChan)
				return
			}
			if t.turn == "even" && t.num%2 == 0 {
				fmt.Println(t.num)
				turnChan <- Turn{num: t.num + 1, turn: "odd"}
			} else {
				// Pass it to the other goroutine
				turnChan <- t
			}
		}
	}()

	// Start the sequence
	turnChan <- Turn{num: 1, turn: "odd"}

	// Wait for both goroutines to finish
	wg.Wait()
}
