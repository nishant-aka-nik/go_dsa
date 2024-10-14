package main

import (
	"fmt"
	"sync"
)

// Input:
// A single integer N which represents the upper limit for finding prime numbers (inclusive). N>=2
// Output:
// A list of all prime numbers up to N.
// Concurrency:
// The program should use multiple Go routines to divide the work of checking for prime numbers.
// Use channels to communicate between Go routines and the main function.
// Ensure that the program handles synchronization properly.
// Example Input: N = 20
// Output: [2, 3, 5, 7, 11, 13, 17, 19]

func main() {
	n := 20

	primePipe := make(chan int)

	var wg sync.WaitGroup

	wg.Add((n - 1))

	for i := 2; i <= n; i++ {
		go IsPrime(i, primePipe, &wg)
	}

	go func() {
		wg.Wait()
		close(primePipe)
	}()

	for value := range primePipe {
		if value != 1 {
			fmt.Println(value)
		}
	}

}

func IsPrime(number int, primePipe chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= number; i++ {
		if number%i == 0 {
			if i == number {
				primePipe <- number
				return
			} else {
				return
			}
		}
	}
}
