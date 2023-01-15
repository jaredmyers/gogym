package main

// Input: strs = ["flower", "flow", "flight"]
// Output: "fl"

// LEETCODE: 14

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for _, val := range strs {
		if len(shortest) > len(val) {
			shortest = val
		}
	}

	prefix := shortest
	flag := false
	for i := 0; i < len(shortest); i++ {
		for _, val := range strs {
			if shortest[i] != val[i] {
				prefix = shortest[:i]
				flag = true
				break
			}
		}
		if flag != false {
			break
		}
	}
	return prefix
}
