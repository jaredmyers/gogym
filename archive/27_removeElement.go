package main

// remove all occurences of val in nums in-place

// LEETCODE: 27

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := i; i < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}
