package main

// LEETCODE: 2108

// time: O(n), space: O(1)
func firstPalindrome(words []string) string {

	for _, word := range words {
		n := len(word)
		flag := true
		for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
			if word[i] != word[j] {
				flag = false
				break
			}
		}
		if flag {
			return word
		}
	}
	return ""
}
