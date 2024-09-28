// https://www.educative.io/courses/concurrency-in-go/deadlock
package main

import "fmt"

// problem code it has deadlock
func main() {
	//create a  unbuffered channel
	// meaning this channel cannot hold any value
	// if value is passed to it value has to be recieved from other code
	deadlockChannel := make(chan int)

	value := <-deadlockChannel // code execution will be blocked here as it is waiting for someone to send value to it but go compiler didnt find it and it throws error
	fmt.Printf("value: %#v\n", value)

}

// //solution to above problem using buffered channel
// import (
// 	"fmt"
// )

// func main() {
// 	goChannel := make(chan int, 1) // buffered channel with capacity of 1

// 	goChannel <- 42 // this will not block since there's space in the buffer

// 	// Receive from the channel
// 	value := <-goChannel
// 	fmt.Println("Received value:", value)
// }

// // solution 2 to above problem using go routine channel
// func main() {
// 	channel := make(chan int)

// 	go func() {
// 		channel <- 42
// 	}()

// 	value := <-channel
// 	fmt.Printf("value: %#v\n", value)
// }

// // solution 2 variation 2 we are passing value to the anonymous function
// func main() {
// 	goChannel := make(chan int)

// 	// Initialize the value to send
// 	valueToSend := 42 // Initialize the variable

// 	// Start a goroutine to send a value
// 	go func(v int) {
// 		goChannel <- v // Send the value passed as a parameter to the channel
// 	}(valueToSend) // Pass the initialized value when calling the anonymous function

// 	// Receive from the channel
// 	value := <-goChannel
// 	fmt.Println("Received value:", value)
// }
