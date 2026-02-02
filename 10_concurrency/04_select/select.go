package main

import (
	"fmt"
	"time"
)

func main() {
    chan1 := make(chan string)
    chan2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        chan1 <- "Data from Channel 1"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        chan2 <- "Data from Channel 2"
    }()

    fmt.Println("--- TOPIC: SELECT ---")

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-chan1:
            fmt.Println("Received:", msg1)
        case msg2 := <-chan2:
            fmt.Println("Received:", msg2)
        }
    }
}