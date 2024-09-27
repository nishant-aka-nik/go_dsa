//https://chatgpt.com/share/66f6568e-afd4-8006-8eb2-126f977e2b5b
package main

import (
	"fmt"
	"sync"
)

// Logger is the singleton struct
type Logger struct {
	// You can add fields like log level, file path, etc.
}

// Instance holds the singleton instance
var instance *Logger

// once is used to ensure the singleton is created only once
var once sync.Once

// GetInstance provides global access to the singleton instance
func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
		// Initialize logger here (e.g., open file, set log level)
		fmt.Println("Logger initialized")
	})
	return instance
}

// Log is a method to demonstrate logging
func (l *Logger) Log(message string) {
	fmt.Println("Log:", message)
}

func main() {
	// Attempt to get multiple instances of Logger
	logger1 := GetInstance()
	logger1.Log("First message")

	logger2 := GetInstance()
	logger2.Log("Second message")

	// Check if both instances are the same
	if logger1 == logger2 {
		fmt.Println("Both logger instances are the same")
	} else {
		fmt.Println("Logger instances are different")
	}
}