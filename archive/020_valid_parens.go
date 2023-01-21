package main

// Input: s = "()"
// Output: true

// Input: s = "()[]{}"
// Output: true

// Input: s = "(]"
// Output: false

// LEETCODE: 20

func isValid(s string) bool {

	brackets := map[string]string{"(": ")", "[": "]", "{": "}"}
	stack := []string{}

	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, strings(s[i]))
			continue
		}

		if string(s[i]) == brackets[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, strings(s[i]))
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
