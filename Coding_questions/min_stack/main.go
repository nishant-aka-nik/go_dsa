package main

import (
	"fmt"
)

// MinStack defines the stack structure
type MinStack struct {
	stack    []int // main stack to store all values
	minStack []int // auxiliary stack to keep track of minimums
}

// Constructor initializes the MinStack
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push adds an element to the stack
func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val) // Push to main stack

	// If minStack is empty or new val is <= current min, push it to minStack
	if len(ms.minStack) == 0 || val <= ms.getCurrentMin() {
		ms.minStack = append(ms.minStack, val)
	}
}

// Pop removes the top element from the stack
func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return // Stack is empty, nothing to pop
	}

	top := ms.stack[len(ms.stack)-1] // Get the top element
	ms.stack = ms.stack[:len(ms.stack)-1] // Remove it from main stack

	// If the popped element is the current min, remove it from minStack as well
	if top == ms.getCurrentMin() {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

// Top retrieves the top element of the stack
func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return -1 // Or handle as per requirements
	}
	return ms.stack[len(ms.stack)-1]
}

// GetMin retrieves the minimum element in the stack
func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return -1 // Or handle as per requirements
	}
	return ms.getCurrentMin()
}

// Helper function to get current minimum
func (ms *MinStack) getCurrentMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

// Example usage
func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // Outputs: -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // Outputs: 0
	fmt.Println(minStack.GetMin()) // Outputs: -2
}