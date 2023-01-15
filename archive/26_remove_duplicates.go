package main

// remove duplicates from sorted array
// LEETCODE 26

func removeDuplicates(arr []int) int {

	for i := 1; i < len(arr); i++ {

		if arr[i-1] == arr[i] {
			for j := i; j < len(arr)-1; j++ {
				arr[j] = arr[j+1]
			}
			arr = arr[:len(arr)-1]
			i--
		}
	}
	return len(arr)
}
