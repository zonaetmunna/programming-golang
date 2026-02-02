package main

import "fmt"

func main() {

    // --- SCOPE 1: ARRAYS (FIXED SIZE) ---
    {
        // Arrays must have a defined size. They cannot grow.
        var arr1 [3]int = [3]int{10, 20, 30}
        
        // Use [...] to let the compiler count the elements
        arr2 := [...]string{"Red", "Green", "Blue"}

        fmt.Println("1. Arrays (Fixed Size):")
        fmt.Printf("   Values: %v | Type: %T\n", arr1, arr1)
        fmt.Printf("   Values: %v | Type: %T\n\n", arr2, arr2)
    }

    // --- SCOPE 2: SLICES (DYNAMIC SIZE) ---
    {
        // Slices do NOT have a size in the brackets.
        myslice := []string{"Apple", "Banana"}
        
        // append() adds elements and returns a new slice
        myslice = append(myslice, "Cherry", "Date")

        fmt.Println("2. Slices (Flexible Size):")
        fmt.Printf("   Values: %v\n", myslice)
        fmt.Printf("   Len: %d | Cap: %d\n\n", len(myslice), cap(myslice))
    }

    // --- SCOPE 3: THE MAKE FUNCTION ---
    {
        // make([]Type, length, capacity)
        // Useful when you know how much memory to pre-allocate
        numbers := make([]int, 3, 5)
        numbers[0] = 100
        numbers[1] = 200
        numbers[2] = 300

        fmt.Println("3. Creating Slice with make():")
        fmt.Printf("   Values: %v | Len: %d | Cap: %d\n\n", numbers, len(numbers), cap(numbers))
    }

    // --- SCOPE 4: SLICING AN EXISTING ARRAY/SLICE ---
    {
        // Syntax: [low : high]
        // Includes 'low', excludes 'high'
        data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

        sub := data[2:5]   // index 2, 3, 4
        firstHalf := data[:5] // from start to 4
        lastHalf := data[5:]  // from 5 to end

        fmt.Println("4. Slicing Operations:")
        fmt.Printf("   Original: %v\n", data)
        fmt.Printf("   Range [2:5]: %v\n", sub)
        fmt.Printf("   Start to 5: %v\n", firstHalf)
        fmt.Printf("   5 to End: %v\n\n", lastHalf)
    }

    // --- SCOPE 5: MULTI-DIMENSIONAL SLICES ---
    {
        // A slice of slices (like a table or grid)
        board := [][]string{
            {"X", "O", "X"},
            {"O", "X", "O"},
            {"X", "O", "X"},
        }

        fmt.Println("5. Multi-Dimensional:")
        fmt.Printf("   Center Value: %s\n", board[1][1])
        fmt.Println("   Full Board:", board, "\n")
    }

    // --- SCOPE 6: COPYING SLICES ---
    {
        source := []int{1, 2, 3}
        destination := make([]int, len(source))
        
        // copy(to, from)
        copy(destination, source)

        fmt.Println("6. Copying Slices:")
        fmt.Printf("   Source: %v, Destination: %v\n", source, destination)
    }
}