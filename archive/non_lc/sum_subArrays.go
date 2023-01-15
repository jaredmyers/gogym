package main

// brute force
func sumSubArrays(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j <= len(arr); j++ {
			for k := i; k < j; k++ {
				result += arr[k]
			}
		}
	}
	return result
}

// math approach
// calculate the number of occurences of a given number
// then multiple that number by the number of its occurences
// add that to total sum
func sumSubArrays2(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result += arr[i] * (i + 1) * (len(arr) - i)
	}
	return result
}
