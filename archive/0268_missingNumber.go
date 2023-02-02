package main

// LEETCODE: 268

// bit manipulation
// time: O(n), space: O(1)
func missingNumber2(nums []int) int {

	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}

// hashset
// time: O(n), space: O(n)
func missingNumber(nums []int) int {

	tracker := map[int]bool{}

	for _, val := range nums {
		tracker[val] = true
	}

	for i := 0; i < len(nums)+1; i++ {
		_, ok := tracker[i]
		if !ok {
			return i
		}
	}

	return -1
}
