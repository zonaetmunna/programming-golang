package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    // Tell the WaitGroup this task is done when the function finishes
    defer wg.Done()

    fmt.Printf("Worker %d starting...\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d finished!\n", id)
}

func main() {
    fmt.Println("--- TOPIC: WAITGROUPS ---")

    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Add 1 task to the counter
        go worker(i, &wg)
    }

    wg.Wait() // Wait here until the counter reaches 0
    fmt.Println("All workers finished.")
}