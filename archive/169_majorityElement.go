package main

// find majority number in array
// LEEDCODE: 169

func majorityElement(nums []int) int {

	tracker := map[int]int{}

	for i := 0; i < len(nums); i++ {

		_, ok := tracker[nums[i]]
		if !ok {
			tracker[nums[i]] = 1
		} else {
			tracker[nums[i]] += 1
		}
	}
	maxNumber := nums[0]
	for key := range tracker {
		if tracker[key] > tracker[maxNumber] {
			maxNumber = key
		}
	}
	return maxNumber

}
