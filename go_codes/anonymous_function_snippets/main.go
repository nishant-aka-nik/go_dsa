package main

import "fmt"

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

}
