package main

// Input: s = "III"
// Output: 3

// LEETCODE: 13

func romanToInt(s string) int {

	lookup := map[byte]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}
	test := []byte(s)

	final := 0
	for i := 0; i < len(test); i++ {
		if len(test) == 1 {
			final += lookup[test[i]]
			break
		}

		if i != len(test)-1 {
			if lookup[test[i]] < lookup[test[i+1]] {
				final += (lookup[test[i+1]] - lookup[test[i]])
				i++
				continue
			}
		}
		final += lookup[test[i]]
	}
	return final
}
