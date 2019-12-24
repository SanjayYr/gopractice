package main

import "fmt"

func main() {
        fmt.Printf("Hello, Sanjay!\n")

    var a int = 10

    /* check the boolean condition using if statement */
    if( a < 20 ) {
      /* if condition is true then print the following */
      fmt.Printf("a is less than 20\n" )
    }
    fmt.Printf("value of a is : %d\n", a)

    if( a < 20 ) {
      /* if condition is true then print the following */
      fmt.Printf("a is less than 20\n" );
    } else {
      /* if condition is false then print the following */
      fmt.Printf("a is not less than 20\n" );
    }

    if( a == 10 ) {
      /* if condition is true then print the following */
      fmt.Printf("Value of a is 10\n" )
    } else if( a == 20 ) {
      /* if else if condition is true */
      fmt.Printf("Value of a is 20\n" )
    } else if( a == 30 ) {
      /* if else if condition is true  */
      fmt.Printf("Value of a is 30\n" )
    } else {
      /* if none of the conditions is true */
      fmt.Printf("None of the values is matching\n" )
    }

    // Switch, expression switch and type switch
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"
   }
   switch {
      case grade == "A" :
         fmt.Printf("Excellent!\n" )
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )
      case grade == "D" :
         fmt.Printf("You passed\n" )
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }

   // Select, channel
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }


}
