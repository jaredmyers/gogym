// keywords: anagram, sort interface, sortrunes,
package main

import "sort"

// LEETCODE: 0049
// Medium

// time: O(nk), space: O(nk), n=len(strs), k=maxLen(str)
func groupAnagrams(strs []string) [][]string {

	n := len(strs)
	preResult := map[string][]string{}
	finalResult := [][]string{}

	for i := 0; i < n; i++ {
		ss := SortString(strs[i])
		_, ok := preResult[ss]
		if !ok {
			preResult[ss] = []string{strs[i]}
		} else {
			preResult[ss] = append(preResult[ss], strs[i])
		}
	}

	for _, val := range preResult {
		finalResult = append(finalResult, val)
	}

	return finalResult
}

// implement sort interface
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Len() int {
	return len(s)
}
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// first try, flawed premise.
func groupAnagrams1(strs []string) [][]string {

	n := len(strs)
	result := [][]string{}

	for i := 0; i < n; i++ {
		result = append(result, []string{strs[i]})
		total1 := 0

		for _, val := range strs[i] {
			total1 += int(val)
		}
		for j := i + 1; j < n; j++ {
			total2 := 0
			for _, val := range strs[j] {
				total2 += int(val)
			}

			if total1 == total2 {
				result[i] = append(result[i], strs[j])
				strs = append(strs[:j], strs[j+1:]...)
				n--
				j--
			}
		}
	}

	return result
}
