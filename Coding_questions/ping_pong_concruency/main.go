package main

import (
	"fmt"
	"sync"
)

type PingPong struct {
	Count int
	Value string
}

func main() {
	limit := 10
	var wg sync.WaitGroup
	wg.Add(2)

	pipe := make(chan PingPong)

	go PingPrinter(pipe, limit, &wg)
	go PongPrinter(pipe, limit, &wg)

	pipe <- PingPong{Count: 0, Value: "Ping"}

	wg.Wait()
}

func PingPrinter(pipe chan PingPong, limit int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range pipe {
		fmt.Println(value)

		if value.Count == limit {
			close(pipe)
			return
		}

		count := value.Count + 1

		pipe <- PingPong{
			Count: count,
			Value: "Ping",
		}

	}
}

func PongPrinter(pipe chan PingPong, limit int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range pipe {
		fmt.Println(value)

		if value.Count == limit {
			close(pipe)
			return
		}

		count := value.Count + 1

		pipe <- PingPong{
			Count: count,
			Value: "Pong",
		}

	}
}


//LEARN in below code i is used an unncessarily and it made the logic wrong fro printing 10 times as it is printing it 20 times
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type PingPong struct {
// 	Count int
// 	Value string
// }

// func main() {
// 	limit := 10
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	pipe := make(chan PingPong)

// 	go PingPrinter(pipe, limit, &wg)
// 	go PongPrinter(pipe, limit, &wg)

// 	pipe <- PingPong{Count: 0, Value: "Ping"}

// 	wg.Wait()
// }

// func PingPrinter(pipe chan PingPong, limit int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	i := 0
// 	for value := range pipe {
// 		fmt.Println(value.Value)
		
// 		if i >= limit {
// 			close(pipe)
// 			return
// 		}

// 		pipe <- PingPong{
// 			Count: i,
// 			Value: "Ping",
// 		}

// 		i++
// 	}
// }

// func PongPrinter(pipe chan PingPong, limit int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	i := 0
// 	for value := range pipe {
// 		fmt.Println(value.Value)
		
// 		if i >= limit {
// 			close(pipe)
// 			return
// 		}

// 		pipe <- PingPong{
// 			Count: i,
// 			Value: "Pong",
// 		}

// 		i++
// 	}
// }