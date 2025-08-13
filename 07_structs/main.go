package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Alice", 30}
    fmt.Println(p.Name, "is", p.Age, "years old")
}
