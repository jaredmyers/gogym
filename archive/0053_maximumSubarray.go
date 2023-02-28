package main

// LEETCODE: 0053
// Medium

import "math"

// time: O(n), space: O(1)
func maxSubArray(nums []int) int {
	n := len(nums)
	current, max := nums[0], nums[0]

	for i := 1; i < n; i++ {
		num := nums[i]
		current = int(math.Max(float64(num), float64(current+num)))
		max = int(math.Max(float64(max), float64(current)))
	}

	return max
}
