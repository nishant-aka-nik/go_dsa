package main

import (
	"fmt"
	"runtime"
	"sync"
)

// // this code is not thread safe it has race issue thats why it gives always random answer
// func main() {

// 	runtime.GOMAXPROCS(4)
// 	value, x := 0, 5
// 	var wg sync.WaitGroup

// 	decrement := func() {
// 		defer wg.Done()
// 		value -= x
// 	}

// 	increment := func() {
// 		defer wg.Done()
// 		value += x
// 	}

// 	for i := 0; i < 200; i++ {
// 		wg.Add(2)
// 		go increment()
// 		go decrement()
// 	}

// 	wg.Wait()
// 	fmt.Println(value)

// }


//------------------------------------------------------------------------------------------
// threadsafe solution - https://www.educative.io/courses/concurrency-in-go/mutex
func main() {

	runtime.GOMAXPROCS(4)
	value, x := 0, 5
	var wg sync.WaitGroup
	var mu sync.Mutex

	decrement := func() {
		defer wg.Done()
		mu.Lock()
		value -= x
		mu.Unlock()
	}

	increment := func() {
		defer wg.Done()
		mu.Lock()
		value += x
		mu.Unlock()
	}

	for i := 0; i < 200; i++ {
		wg.Add(2)
		go increment()
		go decrement()
	}

	wg.Wait()
	fmt.Println(value)

}
