package main

import "fmt"

func main() {
        fmt.Printf("Hello, Sanjay!\n")

    // var balance [10] float32
    // var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

    balance[4] = 50.0
    var salary float32 = balance[3]

    fmt.Println(salary)

    var n [10]int /* n is an array of 10 integers */
   var i,j int

   /* initialize elements of array n to 0 */
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* set element at location i to i + 100 */
   }

   /* output each array element's value */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }

   // multi-dimensional arrays
   // var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
   // var threedim [5][10][4]int

   // a = [3][4]int{
   // {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
   // {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
   // {8, 9, 10, 11}   /*  initializers for row indexed by 2 */


}
