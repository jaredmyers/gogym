package main

// LEETCODE: 217

func containsDuplicate(nums []int) bool {

	tracker := map[int]int{}

	for _, val := range nums {
		_, ok := tracker[val]
		if ok {
			return true
		}

		tracker[val] = 1
	}
	return false
}
