package main

// LEETCODE: 338

// time: O(n log n), space: O(1)
func countBits(n int) []int {
	final := []int{}
	for i := 0; i <= n; i++ {

		popCount := func(x int) int {
			count := 0
			for ; x != 0; count++ {
				x &= x - 1
			}
			return count
		}
		final = append(final, popCount(i))
	}
	return final
}
