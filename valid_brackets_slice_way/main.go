// https://www.educative.io/courses/data-structures-and-algorithms-go/challenge-balanced-parentheses
package main

import "fmt"

var brackets string = "{(})"

// var brackets string = "[()]{}{[()()]()}"

// logic
// if open bracket then push to array
// if close then pop from the array
// and check if the popped bracket is opening complement of closing bracket

func main() {
	var stack []rune

	for _, bracket := range brackets {
		switch bracket {
		case '{':
			stack = append(stack, bracket)
		case '(':
			stack = append(stack, bracket)
		case '[':
			stack = append(stack, bracket)
		case '}':
			lengthOfStack := len(stack)
			lastBraketInStack := stack[lengthOfStack-1]
			if lastBraketInStack != '{' {
				fmt.Println("invalid pair")
				return
			}

			//popping
			stack = stack[:lengthOfStack-1]
		case ')':
			lastBraketInStack := stack[len(stack)-1]
			if lastBraketInStack != '(' {
				fmt.Println("invalid pair")
				return
			}

			//popping
			stack = stack[:len(stack)-1]
		case ']':
			lastBraketInStack := stack[len(stack)-1]
			if lastBraketInStack != '[' {
				fmt.Println("invalid pair")
				return
			}

			//popping
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		fmt.Println("valid pair")
	} else {
		fmt.Println("invalid pair")
	}

}
