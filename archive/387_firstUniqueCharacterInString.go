package main

// LEEDCODE: 387

// Hashmap approach
// time: O(n)  space: O(1) i think
func firstUniqChar(s string) int {

	tracker := map[byte]int{}

	for i := 0; i < len(s); i++ {
		_, ok := tracker[s[i]]
		if !ok {
			tracker[s[i]] = 1
		} else {
			tracker[s[i]] += 1
		}
	}

	for i, val := range s {
		if tracker[byte(val)] == 1 {
			return i
		}
	}
	return -1
}
