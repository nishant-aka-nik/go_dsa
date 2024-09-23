package main

import "fmt"

// // dont use it like this v1
// type Stack []rune

var brackets string = "{(})"

type Stack []rune

func main() {
	//wrong approaches v1
	// stack := new(Stack)
	// stack = append(stack, 'd')

	// //wrong way to make a slice using make we need to give length it is compulsory to give length
	// stack := make([]rune)

	// // ye approch bhi sahi ni hai kyuki agar hume method banani hogi to hum ni bana sakte
	// // interviewer ko code writing style dikhana hoga
	// stack := make([]rune, 0)
	// fmt.Printf("stack: %#v\n", stack)

	var stack Stack

	//to ab stack banane ke baad hume loop krna hai brackets string pe
	// to jab opening bracket ayega to push krna hai stack me
	// opening bracket hamesha push hoga
	// jab closing ayega to - do cheeze check krni hai
	// - kya stack ke top pe opening bracket hai
	//		- agar hai to pop kr de
	// 		- agar ni hai to false return kr de

	// //this has issue as we are not checking the popped pair it showed valid result for invalid pair - {(})
	// // the implementation needs to slightly tweaked
	// for _, char := range brackets {
	// 	switch char {
	// 	case '{', '[', '(':
	// 		stack.Push(char)
	// 	case '}', ']', ')':
	// 		poppedValue := stack.Pop()

	// 		if poppedValue == '}' || poppedValue == ')' || poppedValue == ']' {
	// 			fmt.Println("Invalid brackets")
	// 			return
	// 		}

	// 	}
	// }

	for _, bracket := range brackets {
		switch bracket {
		case '{':
			stack.Push(bracket)
		case '(':
			stack.Push(bracket)
		case '[':
			stack.Push(bracket)
		case '}':
			poppedBracket := stack.Pop()
			if poppedBracket != '{' {
				fmt.Println("Invalid brackets")
				return
			}
		case ')':
			poppedBracket := stack.Pop()
			if poppedBracket != '(' {
				fmt.Println("Invalid brackets")
				return
			}
		case ']':
			poppedBracket := stack.Pop()
			if poppedBracket != '[' {
				fmt.Println("Invalid brackets")
				return
			}
		}

	}

	if len(stack) == 0 {
		fmt.Println("valid brackets")
	} else {
		fmt.Println("Invalid brackets")
	}

}

func (s *Stack) Push(data rune) {
	// learning : yahan jaise s ek pointer bole to address hai usko pehle dereference krre hai phir append operation krre hai
	*s = append(*s, data)
}

func (s *Stack) Pop() rune {
	lengthOfStack := len(*s)

	top := (*s)[lengthOfStack-1]

	*s = (*s)[:lengthOfStack-1]

	return top
}
