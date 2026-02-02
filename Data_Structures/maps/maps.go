package main

import "fmt"

func main() {
    // 1. Initializing a map using make
    // map[KeyType]ValueType
    userAges := make(map[string]int)

    // 2. Assigning values
    userAges["Alice"] = 25
    userAges["Bob"] = 30

    // 3. Map Literal (Short declaration with values)
    
    fruitPrices := map[string]float64{
        "Apple":  0.99,
        "Banana": 0.50,
    }

    fmt.Println("--- TOPIC: MAPS ---")
    fmt.Println("User Ages:", userAges)
    fmt.Printf("Apple Price: $%.2f\n", fruitPrices["Apple"])

    // 4. Checking if a key exists (The 'ok' idiom)
    price, ok := fruitPrices["Orange"]
    if ok {
        fmt.Println("Orange price:", price)
    } else {
        fmt.Println("Orange not found!")
    }

    // 5. Deleting from a map
    delete(fruitPrices, "Banana")

    // 6. Looping over a map
    fmt.Println("\nIterating Map:")
    for key, val := range userAges {
        fmt.Printf("Name: %s, Age: %d\n", key, val)
    }
}