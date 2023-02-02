package main

// LEEDCODE: 5
// Medium

import (
	"fmt"
	"math"
)

// time: O(n^2), space: O(1)
func longestPalindrome(s string) string {

	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))

		if len > (end - start) {
			start = i - (len-1)/2
			end = i + len/2
		}
		fmt.Println("start:", start)
		fmt.Println("  end:", end)
		fmt.Println(" ")
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {

	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	fmt.Println("left:", l, "right:", r)

	return r - l - 1
}
