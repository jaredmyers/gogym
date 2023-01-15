package main

// binary search example

// given sorted array of disinct ints and a target val, return the index of target if found
// if not found, return index of where it would be
// LEETCODE: 35

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(arr []int, target, start, end int) int {

	if end < start {
		return start
	}

	pivotIdx := (start + end) / 2
	if arr[pivotIdx] == target {
		return pivotIdx
	}

	if target < arr[pivotIdx] {
		return binarySearch(arr, target, start, pivotIdx-1)
	} else {
		return binarySearch(arr, target, pivotIdx+1, end)
	}
}
