package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
    "time"
)

// TokenBucket represents the state of the bucket.
type TokenBucket struct {
    capacity       int
    tokens         int
    refillRate     int
    refillInterval time.Duration
    mutex          sync.Mutex
}

// NewTokenBucket initializes a new TokenBucket.
func NewTokenBucket(capacity int, refillRate int, refillInterval time.Duration) *TokenBucket {
    tb := &TokenBucket{
        capacity:       capacity,
        tokens:         capacity, // Start full
        refillRate:     refillRate,
        refillInterval: refillInterval,
    }

    // Start the refill process
    go tb.startRefilling()

    return tb
}

// startRefilling adds tokens to the bucket at regular intervals.
func (tb *TokenBucket) startRefilling() {
    ticker := time.NewTicker(tb.refillInterval)
    for {
        <-ticker.C
        tb.mutex.Lock()
        tb.tokens += tb.refillRate
        if tb.tokens > tb.capacity {
            tb.tokens = tb.capacity
        }
        tb.mutex.Unlock()
    }
}

// Allow checks if a request can be processed.
// It returns true if a token was consumed, false otherwise.
func (tb *TokenBucket) Allow() bool {
    tb.mutex.Lock()
    defer tb.mutex.Unlock()

    if tb.tokens > 0 {
        tb.tokens--
        return true
    }
    return false
}

func main() {
    // Define the rate limiter parameters
    capacity := 10                // Maximum 10 tokens in the bucket
    refillRate := 1               // Add 1 token each refill
    refillInterval := time.Second // Refill every 1 second

    // Initialize the token bucket
    limiter := NewTokenBucket(capacity, refillRate, refillInterval)

    // Define the HTTP handler with rate limiting
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if limiter.Allow() {
            // If allowed, process the request
            fmt.Fprintf(w, "Request processed successfully!\n")
        } else {
            // If not allowed, respond with 429 Too Many Requests
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
        }
    })

    // Start the HTTP server
    port := ":8080"
    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}