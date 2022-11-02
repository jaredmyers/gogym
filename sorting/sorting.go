package sorting

func BubbleSort(arr []int) {
   for i:=0; i<len(arr)-1; i++ {
      for j:=1; j<len(arr)-i; j++ {
         if (arr[j-1] > arr[j]){
            arr[j-1], arr[j] = arr[j], arr[j-1]	
         }
      }
   }
}

func SelectionSort(arr []int) {
   for i:=0; i<len(arr)-1; i++ {
      for j:= i+1; j<len(arr); j++ {
         if arr[j] < arr[i] {
            arr[j], arr[i] = arr[i], arr[j]		
         }
      }
   }
}

func InsertionSort(arr []int) {
   for i:=1; i<len(arr); i++{
      temp := arr[i]
      j := i-1
      for ; j>=0 && arr[j] > temp; j--{
         arr[j+1] = arr[j]
      }
      arr[j+1] = temp
   }
}



