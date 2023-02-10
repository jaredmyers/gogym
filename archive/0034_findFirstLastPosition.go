package main

// LEETCODE: 0034

// O(log n), O(1)
func searchRange2(nums []int, target int) []int {

	first := findBound(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findBound(nums, target, false)
	return []int{first, last}

}

func findBound(nums []int, target int, isFirst bool) int {

	n := len(nums)
	start, end := 0, n-1

	for start <= end {
		pivotIdx := start + (end-start)/2
		if nums[pivotIdx] == target {
			if isFirst {
				if pivotIdx == start || nums[pivotIdx-1] != target {
					return pivotIdx
				}
				end = pivotIdx - 1
			} else {
				if pivotIdx == end || nums[pivotIdx+1] != target {
					return pivotIdx
				}
				start = pivotIdx + 1
			}
		} else if nums[pivotIdx] > target {
			end = pivotIdx - 1
		} else {
			start = pivotIdx - 1
		}
	}
	return -1
}

// first go
// time: O(n) (due to bidirection scan), space: O()
func searchRange(nums []int, target int) []int {
	idx := binarySearch(nums, target, 0, len(nums)-1)
	return fanout(nums, idx)
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

func fanout(nums []int, idx int) []int {

	if idx == -1 {
		return []int{-1, -1}
	}

	start, end := idx, idx

	p1 := idx - 1
	p2 := idx + 1

	for p1 >= 0 {
		if nums[p1] == nums[idx] {
			start = p1
			p1--
			continue
		}
		break
	}

	for p2 < len(nums) {
		if nums[p1] == nums[idx] {
			end = p2
			p2++
			continue
		}
		break
	}
	return []int{start, end}
}
