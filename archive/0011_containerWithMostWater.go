package main

// LEETCODE: 11
// Medium

import "math"

// time: O(n), space: O(1)
func maxArea(height []int) int {

	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		maxArea = int(math.Max(float64(maxArea), math.Min(float64(height[left]), float64(height[right]))*float64(width)))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
