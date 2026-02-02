package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Worker", id, "started")
    time.Sleep(1 * time.Second)
    fmt.Println("Worker", id, "done")
}
func main() {
    fmt.Println("--- TOPIC: WORKERS ---")
    var wg sync.WaitGroup

    numWorkers := 5
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished.")
}
