// golang

package main

import (
		"fmt"
		sort "github.com/jaredmyers/scratch/sorting"
)

func main() {
		
		arr := []int{88,234,333,78,23,1,20}
		fmt.Println(arr)
		sort.BubbleSort(arr)
		fmt.Println(arr)
}
