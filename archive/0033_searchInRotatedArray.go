package main

// LEETCODE: 0033

// time: O(log n), space: O(1)
func search2(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		pivotIdx := start + (end-start)/2
		if nums[pivotIdx] == target {
			return pivotIdx
		}

		if nums[pivotIdx] >= nums[start] {
			if target >= nums[start] && target < nums[pivotIdx] {
				end = pivotIdx - 1
			} else {
				start = pivotIdx + 1
			}
		} else {
			if target <= nums[end] && target > nums[pivotIdx] {
				start = pivotIdx + 1
			} else {
				end = pivotIdx - 1
			}
		}
	}
	return -1
}

// first go
// time: O(log n), O(1)
func search(nums []int, target int) int {

	n := len(nums)
	pivot := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			pivot = i
		}
	}

	result := 0
	tmp := pivot
	if pivot == 0 {
		tmp = 1
	}
	if target >= nums[0] && target <= nums[tmp-1] {
		result = binarySearch(nums, target, 0, pivot)
	} else {
		result = binarySearch(nums, target, pivot, n-1)
	}

	return result

}

func binarySearch(nums []int, target, start, end int) int {

	if end < start {
		return -1
	}

	pivotIdx := start + (end-start)/2
	if nums[pivotIdx] == target {
		return pivotIdx
	}

	if target < nums[pivotIdx] {
		return binarySearch(nums, target, start, pivotIdx-1)
	}
	return binarySearch(nums, target, pivotIdx+1, end)
}
