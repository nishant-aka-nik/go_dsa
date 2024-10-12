package main

import (
    "fmt"
    "time"
)

func ping(pongChan chan string) {
    for i := 0; i < 5; i++ {
        time.Sleep(500 * time.Millisecond) // Simulate work
        pongChan <- "ping" // Send "ping" to pong channel
        fmt.Println("Sent: ping")
    }
    close(pongChan) // Close the channel when done
}

func pong(pongChan chan string) {
    for msg := range pongChan {
        fmt.Println("Received:", msg)
        time.Sleep(500 * time.Millisecond) // Simulate work
        fmt.Println("Sent: pong")
    }
}

func main() {
    pongChan := make(chan string)

    go ping(pongChan) // Start the ping goroutine
    go pong(pongChan) // Start the pong goroutine

    // Wait for goroutines to finish
    time.Sleep(6 * time.Second) // Adjust sleep time as needed
    fmt.Println("Finished")
}