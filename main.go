// golang

package main

import (
	"fmt"
	sort "github.com/jaredmyers/scratch/sorting"
)

func main() {

	arr := []int{897, 12345, 233, 123, 78, 63, 1, 23}
	fmt.Println(arr)
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
