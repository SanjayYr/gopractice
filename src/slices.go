package main

import "fmt"

func main() {
   var numbers1 = make([]int,3,5)
   printSlice(numbers1)

   var numbers2 []int
   printSlice(numbers2)

   if(numbers2 == nil){
      fmt.Printf("slice is nil\n")
   }

   // sub-slicing
   /* create a slice */
   numbers := []int{0,1,2,3,4,5,6,7,8}
   printSlice(numbers)

   /* print the original slice */
   fmt.Println("numbers ==", numbers)

   /* print the sub slice starting from index 1(included) to index 4(excluded)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* missing lower bound implies 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* missing upper bound implies len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers3 := make([]int,0,5)
   printSlice(numbers3)

   /* print the sub slice starting from index 0(included) to index 2(excluded) */
   number2 := numbers[:2]
   printSlice(number2)

   /* print the sub slice starting from index 2(included) to index 5(excluded) */
   number3 := numbers[2:5]
   printSlice(number3)

   // append() and copy()

   var numbers5 []int
   printSlice(numbers5)

   /* append allows nil slice */
   numbers5 = append(numbers5, 0)
   printSlice(numbers5)

   /* add one element to slice*/
   numbers5 = append(numbers5, 1)
   printSlice(numbers5)

   /* add more than one element at a time*/
   numbers5 = append(numbers5, 2,3,4)
   printSlice(numbers5)

   /* create a slice numbers1 with double the capacity of earlier slice*/
   numbers6 := make([]int, len(numbers5), (cap(numbers5))*2)

   /* copy content of numbers to numbers1 */
   copy(numbers6,numbers5)
   printSlice(numbers6)
}
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
