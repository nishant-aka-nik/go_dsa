package main

import (
	"fmt"
	"sync"
)

// 2 gr

// 1 gr generate 0 to 5
// 2 gr consumer take from producer

func main() {
	pipe := make(chan int)

	var wg1 sync.WaitGroup

	var wg2 sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go Generator(pipe, &wg1)
	}

	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go Consumer(pipe, &wg2)
	}

	go func(){
		wg1.Wait()
		close(pipe)
	}()

	wg2.Wait()
}

func Generator(pipe chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		pipe <- i
	}
}

func Consumer(pipe chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range pipe {
		fmt.Println("Value from generator ", value)
	}

}
