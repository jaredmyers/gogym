package main

// LEETCODE: 1351

// time: O(n), space: O(1)
func countNegatives(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])
	total := 0

	for i := 0; i < n; i++ {

		rowTotal := 0
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] < 0 {
				rowTotal++
			}
		}
		total += rowTotal
	}
	return total
}
