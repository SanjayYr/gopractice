package main

import "fmt"

func main() {
        fmt.Printf("Hello, Sanjay!\n")

    // Operators
    // Arithmetic
    p, q := 13, 14
    fmt.Printf("\nArithmetic\n")
    fmt.Println(p+q)
    fmt.Println(p-q)
    fmt.Println(p*q)
    fmt.Println(p/q)
    fmt.Println(p%q)
    // increment and decrement operators are statements, not values
    p++
    q--
    fmt.Printf("increment p: %d\n", p)
    fmt.Printf("decrement q: %d\n", q)

    // Relational
    fmt.Printf("\nRelational\n")
    fmt.Println(p==q)
    fmt.Println(p!=q)
    fmt.Println(p>q)
    fmt.Println(p<q)
    fmt.Println(p>=q)
    fmt.Println(p<=q)

    // Logical
    fmt.Printf("\nLogical\n")
    fmt.Println((p!=0) && (q!=0))
    fmt.Println((p!=0) || (q!=0))
    fmt.Println(!(p!=0))

    // Bitwise
    fmt.Printf("\nBitwise\n")

    fmt.Println(p&q)
    fmt.Println(p|q)
    fmt.Println(p^q)
    fmt.Println(p<<2)
    fmt.Println(p>>2)

    // Assignment
    // = , +=, -=, *=, /=, %=, <<==, >>==, &=, |=, ^=

    // &, *

}
