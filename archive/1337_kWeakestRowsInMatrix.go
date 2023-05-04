package main

// LEETCODE: 1337

import "sort"

// time: O(m(n + log m)), space: O(m)
func kWeakestRows(mat [][]int, k int) []int {

	tracker := map[int][]int{}
	// time: O(nm)
	for i := 0; i < len(mat); i++ {
		temp := 0
		for j := 0; j < len(mat[0]); j++ {
			temp += mat[i][j]
		}

		// time: O(1)
		_, ok := tracker[temp]
		if ok {
			tracker[temp] = append(tracker[temp], i)
		} else {
			tracker[temp] = []int{i}
		}
	}

	// time: O(m log m)
	keys := []int{}
	for k := range tracker {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// time: O(keys)
	result := []int{}
	for _, v1 := range keys {
		for _, v2 := range tracker[v1] {
			result = append(result, v2)
			if len(result) == k {
				break
			}
		}
		if len(result) == k {
			break
		}
	}
	return result
}
