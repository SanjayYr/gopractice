package main

import (
    "fmt"
    "math"
)

func main() {
        fmt.Printf("Hello, Sanjay!\n")

    // max
    fmt.Println(max(5,6))

    // swap
    a, b := swap("Mahesh", "Kumar")
    fmt.Println(a, b)

    // function as a value
    getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* use the function */
   fmt.Println(getSquareRoot(9))


}

/* function returning the max between two numbers */
func max(num1 int, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

// Returning multiple values
func swap(x, y string) (string, string) {
   return y, x
}

// call by value
func swap2(x int, y int) {
   var temp int
   temp = x    /* save the value at address x */
   x = y    /* put y into x */
   y = temp    /* put temp into y */
}

// call by reference
func swap3(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x */
   *x = *y    /* put y into x */
   *y = temp    /* put temp into y */
}
