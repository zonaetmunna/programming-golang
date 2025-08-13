package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Bangladesh": "Dhaka",
        "Japan":      "Tokyo",
        "USA":        "Washington DC",
    }

    capitals["India"] = "New Delhi"
    delete(capitals, "USA")

    for country, capital := range capitals {
        fmt.Println(country, "→", capital)
    }
}
