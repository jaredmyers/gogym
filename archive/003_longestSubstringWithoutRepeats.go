package main

import "math"

// LEEDCODE: 003
// medium

// time: O(n), space: O(min(m, n))
func lengthOfLongestSubstring(s string) int {

	final := 0
	tracker := map[byte]int{}

	i := 0
	for j := 0; j < len(s); j++ {

		_, ok := tracker[s[j]]
		if ok {
			i = int(math.Max(float64(tracker[s[i]]), float64(j-i+1)))
		}
		final = int(math.Max(float64(final), float64(j-i+1)))
		tracker[s[j]] = j + 1

	}
	return int(final)

}
