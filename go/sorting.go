package main

import (
   "fmt"
)

func main() {

   array := []int{11,14,3,8,18,17,43};
   fmt.Println(bubbleSort(array))
}

func bubbleSort(array []int) []int {
   for i:=0; i<len(array)-1; i++ {
      for j:=1; j<len(array)-i; j++ {
         if (array[j-1] > array[j]){
            array[j-1], array[j] = array[j], array[j-1]	
         }
      }
   }
   return array
}

func selectionSort(arr []int) []int {
   for i:=0; i<len(arr)-1; i++ {
      for j:= i+1; j<len(arr); j++ {
         if arr[j] < arr[i] {
            arr[j], arr[i] = arr[i], arr[j]		
         }
      }
   }
   return arr
}

func insertionSort(arr []int) []int {
   for i:=1; i<len(arr); i++{
      temp := arr[i]
      j := i-1
      for ; j>=0 && arr[j] > temp; j--{
         arr[j+1] = arr[j]
      }
      arr[j+1] = temp
   }
   return arr
}




