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
