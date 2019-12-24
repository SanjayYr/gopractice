package main

import "fmt"

func main() {
        fmt.Printf("Hello, Sanjay!\n")

    // For loop
    var b int = 15
   var a int
   numbers := [6]int{1, 2, 3, 5}

   /* for loop execution */
   for a := 0; a < 10; a++ {
      fmt.Printf("value of a: %d\n", a)
   }
   for a < b {
      a++
      fmt.Printf("value of a: %d\n", a)
   }
   for i,x:= range numbers {
      fmt.Printf("value of x = %d at %d\n", x,i)
   }

   // break
   var c int = 10

   /* for loop execution */
   for c < 20 {
      fmt.Printf("value of c: %d\n", c);
      c++;
      if c > 15 {
         /* terminate the loop using break statement */
         break;
      }
   }


   // continue
   var d int = 10

   /* do loop execution */
   for d < 20 {
      if d == 15 {
         /* skip the iteration */
         d = d + 1;
         continue;
      }
      fmt.Printf("value of d: %d\n", d);
      d++;
   }

   // goto
   var e int = 10

   /* do loop execution */
   LOOP: for e < 20 {
      if e == 15 {
         /* skip the iteration */
         e = e + 1
         goto LOOP
      }
      fmt.Printf("value of e: %d\n", e)
      e++
   }
}
