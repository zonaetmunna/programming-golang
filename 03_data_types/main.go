package main

import "fmt"

func main() {

    // --- SCOPE 1: EXPLICIT DECLARATIONS ---
    {
        var a int = 10
        var b float64 = 3.14
        var c bool = true
        var d string = "Go"
        
        fmt.Println("1. Explicit Style:")
        fmt.Printf("   Types: %T, %T, %T, %T\n", a, b, c, d)
        fmt.Printf("   Values: %v, %v, %v, %s\n\n", a, b, c, d)
    }

    // --- SCOPE 2: INFERRED / SHORT DECLARATIONS ---
    {
        a := 50             // compiler assumes int
        b := 99.9           // compiler assumes float64
        c := false          // compiler assumes bool
        d := "Learning"     // compiler assumes string

        fmt.Println("2. Short/Inferred Style (:=):")
        fmt.Printf("   %T, %T, %T, %T\n\n", a, b, c, d)
    }

    // --- SCOPE 3: ZERO VALUES (DEFAULT INITIALIZATION) ---
    {
        var i int
        var f float64
        var b bool
        var s string

        fmt.Println("3. Zero Values (Defaults):")
        fmt.Printf("   Int: %d | Float: %f | Bool: %t | String: [%s]\n\n", i, f, b, s)
    }

    // --- SCOPE 4: SPECIFIC SIZED NUMERICS ---
    {
        var smallNum int8 = 127
        var bigNum int64 = 9223372036854775807
        var uNum uint = 500 // Only positive numbers
        
        fmt.Println("4. Sized Numerics:")
        fmt.Printf("   int8: %d, int64: %d, uint: %d\n\n", smallNum, bigNum, uNum)
    }

    // --- SCOPE 5: CHARACTER TYPES (RUNES & BYTES) ---
    {
        var letter byte = 'A'     // uint8 (ASCII)
        var emoji rune = '🎯'      // int32 (Unicode/UTF-8)

        fmt.Println("5. Character Types:")
        fmt.Printf("   Byte: %v (Type: %T)\n", letter, letter)
        fmt.Printf("   Rune: %v (Type: %T) - Literal: %c\n\n", emoji, emoji, emoji)
    }

    // --- SCOPE 6: MULTIPLE ASSIGNMENT ---
    {
        name, score, active := "User1", 95, true

        fmt.Println("6. Parallel Assignment:")
        fmt.Printf("   %s has a score of %d. Status: %t\n", name, score, active)
    }
}