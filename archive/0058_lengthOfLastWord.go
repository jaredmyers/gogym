package main

// given a string, return length of last word

func lengthOfLastWord(s string) int {

	start := 0
	end := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 && start == 0 {
			start = i
		}

		if s[i] == 32 && start != 0 {
			end = i
			break
		} else if i == 0 {
			end = i - 1
		}
	}

	if start == 0 && end == 0 {
		return 1
	}

	return (start - end)
}
