package main

import "fmt"

func main() {
        fmt.Printf("Hello, Sanjay!\n")
    var x float64
    x = 20.0
    fmt.Printf("x = %f\n", x)
    fmt.Printf("x is of type %T\n", x)

    // declaration without type, requires initialization during declaration
    y := 30
    fmt.Printf("y = %d\n", y)
    fmt.Printf("y is of type %T\n", y)

    // tuple style declaration
    var a, b, c = 3, 4, "foo"
    d, e, f := 5, 6.0, "bar"

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println()

    // string
    var s = "string variable"

    fmt.Println(s)
    fmt.Println(string(s[0]))
    fmt.Println(string(s[0:5]))
    fmt.Println(string('a'))
    fmt.Println(int('a'))

    // constants
    const LEN int=10
    const WIDTH int=20

    fmt.Printf("area is : %d\n", LEN*WIDTH)

    // Operators
    // Arithmetic
    p, q := 13, 14
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

    // bool
    var flag bool
    flag = true

    fmt.Println(flag)
}
