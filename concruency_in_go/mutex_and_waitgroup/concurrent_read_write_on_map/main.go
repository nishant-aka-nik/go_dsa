//https://www.educative.io/courses/concurrency-in-go/solution-perform-a-concurrent-read-write-operation-on-a-map
package main

import (
	"fmt"
	"sync"
)


// //faulty code as waitgroup is not used in this code to make the main thread wait
// // it will end without waiting for other go routines to complete
// func main() {
// 	var mu sync.Mutex
// 	dataMap := make(map[int]string)

// 	go WriteMap(dataMap, &mu)
// 	go ReadMap(dataMap, &mu)

// }

// //Lock locks m. If the lock is already in use, the calling goroutine blocks until the mutex is available.

// func WriteMap(data map[int]string, mu *sync.Mutex) {
// 	mu.Lock()
// 	data[0] = "hello"
// 	defer mu.Unlock()
// }

//------------------------------------------Incomplete and useless example-----------------------------------------------------------------------------

func main() {
	var mu sync.Mutex
	dataMap := make(map[int]string)

	var wg sync.WaitGroup

	wg.Add(2)

	go WriteMap(dataMap, &mu, &wg)
	go ReadMap(dataMap, &mu, &wg)

	wg.Wait()
}

//Lock locks m. If the lock is already in use, the calling goroutine blocks until the mutex is available.

func WriteMap(data map[int]string, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	data[0] = "hello"
	defer mu.Unlock()
	defer wg.Done()
}

func ReadMap(data map[int]string, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	value := data[0]
	fmt.Printf("value: %#v\n", value)
	defer mu.Unlock()
	defer wg.Wait()
	
}
