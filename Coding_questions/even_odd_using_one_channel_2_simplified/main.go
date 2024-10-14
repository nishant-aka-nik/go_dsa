package main

import (
	"fmt"
	"sync"
)

func main() {
	limit := 10

	pipe := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go EvenPrinter(pipe, &wg, limit)

	go OddPrinter(pipe, &wg, limit)

	pipe <- 0

	wg.Wait()

}

func EvenPrinter(pipe chan int, wg *sync.WaitGroup, limit int) {
	defer wg.Done()

	for value := range pipe {
		if value == limit {
			// ye close hoga Odd wala band hoga
			// ye wala close dono go routine ko terminate kr dega
			// yahan se close karo wahan odd wala for loop terminate ho jaega
			close(pipe)
			return
		}

		if value%2 == 0 {
			fmt.Println("Even ", value)
			value++
			pipe <- value
		} else {
			pipe <- value
		}

	}

}

func OddPrinter(pipe chan int, wg *sync.WaitGroup, limit int) {
	defer wg.Done()

	for value := range pipe {
		if value == limit {
			close(pipe)
			return
		}

		if value%2 != 0 {
			fmt.Println("Odd  ", value)
			value++
			pipe <- value
		} else {
			pipe <- value
		}

	}

}
