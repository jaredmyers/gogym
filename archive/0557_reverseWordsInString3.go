package main

// LEETCODE: 0557

// time: O(n), space: O(1)?
// two pointer
func reverseWords(s string) string {

	x := []byte(s)
	lastSpaceIndex := -1
	for i := 0; i <= len(x); i++ {
		// short circuit
		if i == len(x) || x[i] == 32 {
			startIndex := lastSpaceIndex + 1
			endIndex := i - 1
			for startIndex < endIndex {
				x[startIndex], x[endIndex] = x[endIndex], x[startIndex]
				startIndex++
				endIndex--
			}
			lastSpaceIndex = i
		}
	}
	return string(x)
}

// first go
// time: O(n), space: O(n)
func reverseWords1(s string) string {

	final := []byte{}

	j := 0
	for i := 0; i < len(s); i++ {

		if s[i] == 32 || i == len(s)-1 {
			tmp := j
			if i != len(s)-1 {
				j = i - 1
			} else {
				j = i
			}

			for ; j >= tmp; j-- {
				final = append(final, s[j])
			}
			if i != len(s)-1 {
				final = append(final, 32)
			}
			j = i + 1
		}
	}
	return string(final)
}
