package main

// LEEDCODE: 242

import (
	"sort"
	"strings"
)

// time: O(n), space: O(1)
func isAnagram2(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	s1 := []rune(s)
	s2 := []rune(t)
	tracker := make([]int, 26)

	a := []rune("a")
	for i := 0; i < len(s1); i++ {
		tracker[(s1[i]-a[0])]++
		tracker[(s2[i]-a[0])]--
	}

	for _, val := range tracker {
		if val != 0 {
			return false
		}
	}
	return true
}

// first try, some fun with sort/strings pkgs
// O(n log n)
func isAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	p1 := strings.Split(s, "")
	p2 := strings.Split(t, "")

	sort.Strings(p1)
	sort.Strings(p2)

	x := strings.Join(p1, "")
	y := strings.Join(p2, "")

	if x != y {
		return false
	}
	return true
}
