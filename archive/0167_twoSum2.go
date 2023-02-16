package main

// LEEDCODE: 167
// Medium

// time: O(n), space: O(1)
func twoSum2(numbers []int, target int) []int {

	final := []int{}
	left := 0
	right := len(numbers) - 1

	for left < right {
		temp := numbers[left] + numbers[right]
		if temp == target {
			break
		}

		if temp < target {
			left++
		} else {
			right--
		}
	}

	final = append(final, left+1, right+1)
	return final
}
