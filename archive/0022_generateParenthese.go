package main

// LEETCODE: 0022
// Medium

// time: O(), space: O()
func generateParenthese(n int) []string {

	result := []string{}
	if n == 0 {
		result = append(result, "")
	} else {
		for i := 0; i < n; i++ {
			left := generateParenthese(i)
			for _, lv := range left {
				right := generateParenthese(n - 1 - i)
				for _, rv := range right {
					add := "(" + lv + ")" + rv
					result = append(result, add)
				}
			}
		}
	}
	return result
}
