package main

import (
	"fmt"
)

func Even(evenCh chan int, done chan int) {
	// this is important as it may lead to memory leakage
	defer close(evenCh)

	i := 0
	for {
		select {
		case <-done:
			return
		case evenCh <- i:
			i += 2
		}
	}
}

func Odd(oddCh chan int, done chan int) {
	defer close(oddCh)

	i := 1
	for {
		select {
		case <-done:
			return
		case oddCh <- i:
			i += 2
		}
	}
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	// or advance to ye done to recieve only bana do
	done := make(chan int)

	go Even(evenCh, done)
	go Odd(oddCh, done)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(<-evenCh)
		} else {
			fmt.Println(<-oddCh)
		}
	}

	done <- 0

}
