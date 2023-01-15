// Golang
// testing
package main

// takes in sorted nums1 int slice and m = len(nums1)
// takes in sorted nums2 int slice and n = len(nums2)
// sorts as num1 inplace
func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := len(nums1) - 1; i >= 0; i-- {
		if nums1[i] == 0 {
			nums1 = nums1[:i]
		} else {
			break
		}
	}

	p1 := len(nums1) - 1
	p2 := n - 1

	for range nums2 {
		nums1 = append(nums1, 0)
	}

	for p := len(nums1) - 1; p >= 0; p-- {

		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}
}
