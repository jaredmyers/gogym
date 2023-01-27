package main

// LEEDCODE: 283

// time: O(n), space: O(1)
func moveZeros2(nums []int) {

	for lastNonZero, curr := 0, 0; curr < len(nums); curr++ {
		if nums[curr] != 0 {
			nums[lastNonZero], nums[curr] = nums[curr], nums[lastNonZero]
			lastNonZero++
		}
	}
}

// Time: O(n), Space: O(1)
func moveZeros(nums []int) {

	lastNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}

}
