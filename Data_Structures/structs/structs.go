package main

import "fmt"

// 1. Defining a Struct (The Blueprint)

type Student struct {
    Name  string
    ID    int
    Grade float64
}

func main() {
    // 2. Creating an instance of a struct
    s1 := Student{
        Name:  "Rahim",
        ID:    101,
        Grade: 3.85,
    }

    // 3. Accessing and Modifying Fields
    fmt.Println("--- TOPIC: STRUCTS ---")
    fmt.Println("Student Name:", s1.Name)
    
    s1.Grade = 4.0 // Updating a value
    fmt.Printf("Updated Grade: %.2f\n", s1.Grade)

    // 4. Anonymous Struct (Good for temporary data)
    point := struct {
        X, Y int
    }{
        X: 10,
        Y: 20,
    }

    fmt.Printf("Point: %+v\n", point) // %+v prints field names + values
}