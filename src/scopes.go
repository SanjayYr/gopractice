package main

import "fmt"

/* global variable declaration */
var g int

func main() {
        fmt.Printf("Hello, Sanjay!\n")

    /* local variable declaration */
   var a, b, c int

   /* actual initialization */
   a = 10
   b = 20
   c = a + b

   fmt.Printf ("value of a = %d, b = %d and c = %d\n", a, b, c)


   // global
   g = a + b
   fmt.Printf ("value of a = %d, b = %d and g = %d\n", a, b, g)

   // local g takes precedence
   var g int = 10
   fmt.Printf ("value of a = %d, b = %d and g = %d\n", a, b, g)
}
