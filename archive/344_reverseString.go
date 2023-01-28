package main

// LEETCODE: 344

// two pointers approach
// time: O(n), space: O(1)
func reverseString(s []byte) {

	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
