package main

// LEEDCODE: 350

// hashmap approach
func intersect(nums1, nums2 []int) []int {

	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	tracker := map[int]int{}
	final := []int{}

	tracker = createHash(nums1)
	final = compare(tracker, nums2)
	return final
}

func createHash(arr []int) map[int]int {

	tracker := map[int]int{}
	for i := 0; i < len(arr); i++ {
		_, ok := tracker[arr[i]]
		if !ok {
			tracker[arr[i]] = 1
		} else {
			tracker[arr[i]] += 1
		}
	}
	return tracker
}

func compare(tracker map[int]int, arr []int) []int {

	final := []int{}
	for _, val := range arr {
		_, ok := tracker[val]

		if !ok {
			continue
		}
		final = append(final, val)
		tracker[val] -= 1
		if tracker[val] < 1 {
			delete(tracker, val)
		}
	}
	return final
}
