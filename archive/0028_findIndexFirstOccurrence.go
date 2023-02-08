package main

// LEETCODE: 0028

// time: O(n), space: O(1)
func strStr(haystack, needle string) int {

	n := len(needle)
	for i, j := 0, n; j <= len(haystack); i, j = i+1, j+1 {
		if haystack[i:j] == needle {
			return i
		}
	}
	return -1
}
