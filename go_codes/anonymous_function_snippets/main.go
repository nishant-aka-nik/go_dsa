package main

import (
	"fmt"
	"sync"
)

func main() {
	// We can assign an anonymous function to a variable and use the variable to execute the function.
	anonymousFunc := func(input string) string {
		// Do some stuff
		return input
	}

	x := anonymousFunc("input")
	fmt.Printf("x: %#v\n", x)

	//---------------------------------------------------------------------------------------------------------------------
	// If we put the closing brackets just after declaring the function, it’s invoked immediately.
	func() {
	}()

	//---------------------------------------------------------------------------------------------------------------------
	//We can pass the anonymous function as an input to the regular functions as follows:
	parameter := func(input string) {
		fmt.Println(input)
	}

	func(v string, anonfunction func(v string)) {
		anonfunction(v)

	}("mystring", parameter)

	////---------------------------------------------------------------------------------------------------------------------
	//// this below code will always print 11 we are not passing the value of i in anonymous function
	// var wg sync.WaitGroup
	// for i := 1; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	//----------------------------------------------Fixed code------------------------------------------------------------------
	// this will print integers from 1 to 10 in random order based on go routine execution
	// jo pehle aya vo pehle print krra
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(value int) {
			fmt.Println(value)
			wg.Done()
		}(i)
	}

	wg.Wait()

}
