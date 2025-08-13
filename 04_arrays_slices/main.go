package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    slice := []string{"Go", "Python", "Rust"}

    slice = append(slice, "Java")

    fmt.Println("Array:", arr)
    fmt.Println("Slice:", slice)
    fmt.Println("Third element of slice:", slice[2])
}
