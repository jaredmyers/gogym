package main

import (
		"fmt"
)

func bubbleSort(array[] int) []int {
   for i:=0; i<len(array); i++ {
      for j:=1; j<len(array)-i; j++ {
         if (array[j-1] > array[j]){
		    array[j-1], array[j] = array[j], array[j-1]	
		 }
      }
   }
   return array
}

func main() {
   array := []int{11,14,3,8,18,17,43};
   fmt.Println(len(array))
   fmt.Println(bubbleSort(array))
}


