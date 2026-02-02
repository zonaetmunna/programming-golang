package main

import (
	"fmt"
	"time"
)

func display(message string) {
    for i := 1; i <= 3; i++ {
        fmt.Println(message, ":", i)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    fmt.Println("--- TOPIC: GOROUTINES ---")

    // This runs in a new goroutine (background)
    go display("Async Task")

    // This runs in the main goroutine
    display("Main Task")

    fmt.Println("Done!")
}