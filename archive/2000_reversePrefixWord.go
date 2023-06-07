package main

// LEETCODE: 2000

// time: O(n), space: O(n)
func reversePrefix(word string, ch byte) string {

	pos := 0
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			pos = i
			break
		}
	}

	result := ""
	for i := pos; i >= 0; i-- {
		result += string(word[i])
	}

	for i := pos + 1; i < len(word); i++ {
		result += string(word[i])
	}

	return result
}
