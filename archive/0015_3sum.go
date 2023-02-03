package main

import "sort"

// LEEDCODE: 15
// Medium

// time: O(n^2), space: O(logn)
func threeSum(nums []int) [][]int {
	final := [][]int{}
	// sort first, makes skipping repeats easier
	// & we'll already be at worst O(n^2) time, so it doesn't change our overall cost
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1

		// makes sure to skip repeat nums for i
		if i == 0 || nums[i-1] != nums[i] {

			// holding i, we do two pointer method on i+1 to end
			for left < right {
				temp := nums[i] + nums[left] + nums[right]

				// if 0 we found a triplet
				if temp == 0 {
					final = append(final, []int{nums[i], nums[left], nums[right]})
					left++
					right--
					// skip repeat nums for left
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					continue
				}

				if temp < 0 {
					left++
				} else {
					right--
				}
			}
		}
	}
	return final
}
