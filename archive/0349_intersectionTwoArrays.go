package main

// LEETCODE: 0349

// time: O(n*m), space: O(n*m)
func intersection(nums1, nums2 []int) []int {

	preResult := map[int]struct{}{}
	for _, val := range nums1 {
		for _, val2 := range nums2 {
			if val == val2 {
				_, ok := preResult[val]
				if !ok {
					preResult[val] = struct{}{}
				}
			}
		}
	}

	result := []int{}
	for k, _ := range preResult {
		result = append(result, k)
	}

	return result
}
