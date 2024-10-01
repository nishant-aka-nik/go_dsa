package main

import "fmt"

// //-----------------------------------------------Buggy Code-------------------------------------------------------------
// func main() {
// 	channel := make(chan int) //this is unbuffered channel

// 	//now we want to put value to channel
// 	// use the <- to determine where data is going
// 	// is it going into the channel or out from the channel

// 	channel <- 10

// 	valueOutFromChannel := <-channel

// 	fmt.Println(valueOutFromChannel)

// }

// // This occurs because the task of the channel is to
// // ------------communicate between goroutines and synchronize them--------------
// // ------------The channel only works when the sender and receiver are both ready--------------
// // The code executes in sequential order. Thus, the value is sent, but there is no one to receive it,
// // and the code waits until a corresponding goroutine is ready to accept or receive the value. Thats why the error occurred.

// //-----------------------------------------------Fixed Code-------------------------------------------------------------
func main() {
	goChannel := make(chan int)

	// now what we have to do is to run go routine and pass value to the channel
	// we will use the anonymous function
	// we can also declare the function

	// first lets declare it and then make it run asynchronously
	// since it is a anonymous function it can access he channel (parent variable)
	// //declaration
	// func() {
	// 	goChannel <- 10
	// }()

	//making it go routine to run it asynchronously
	go func() {
		goChannel <- 10
	}()

	valueFromChannel := <-goChannel
	fmt.Printf("valueFromChannel: %#v\n", valueFromChannel)

	// code part 2
	go PutSomeValueToChannel(30, goChannel)

	takeValueFromChannel := <-goChannel

	fmt.Printf("takeValueFromChannel: %#v\n", takeValueFromChannel)
}

func PutSomeValueToChannel(value int, goChannel chan int) {
	goChannel <- value
}

//-----------------------------------------------learn-------------------------------------------------------------
// unbuffered channel needs to send and recieve value at the same time
// •	It’s idiomatic in Go to name channels with a ch suffix to indicate their purpose.
// •	Example: goCh instead of goChannel.