package main

import "fmt"

func sendData(ch chan string) {
    // Sending data into the channel
    ch <- "Hello from the background!"
}

func main() {
    fmt.Println("--- TOPIC: CHANNELS ---")

    // Create a channel: make(chan Type)
    messageChannel := make(chan string)

    // Start goroutine
    go sendData(messageChannel)

    // Receive data from the channel (this waits until data arrives)
    received := <-messageChannel

    fmt.Println("Received:", received)
}