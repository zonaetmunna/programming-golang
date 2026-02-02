package main

import "fmt"

func main() {
    // 1. Basic Variable
    num := 100

    // 2. Creating a Pointer (Store the address of num)
    // '&' operator gets the address
    
    ptr := &num 

    fmt.Println("--- TOPIC: POINTERS ---")
    fmt.Println("Value of num:   ", num)      // 100
    fmt.Println("Address of num: ", ptr)      // e.g., 0xc0000120f0

    // 3. Dereferencing (Access value via address)
    // '*' operator gets value from the pointer
    fmt.Println("Value via ptr:  ", *ptr)    // 100

    // 4. Changing value through the pointer
    *ptr = 200
    fmt.Println("New Value of num:", num)   // num is now 200

    // 5. Pointer Zero Value
    var emptyPtr *int
    fmt.Println("Nil Pointer:    ", emptyPtr) // <nil>
}