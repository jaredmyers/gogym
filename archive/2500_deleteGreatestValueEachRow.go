package main

// LEETCODE: 2500

import (
	"math"
	"sort"
)

// time: O(n), space: O(1)
func deleteGreatestValue(grid [][]int) int {

	for i := range grid {
		sort.Ints(grid[i])
	}

	sum := 0

	for i := range grid[0] {
		maxVal := 0
		for j := range grid {
			maxVal = int(math.Max(float64(grid[j][i]), float64(maxVal)))
		}
		sum += maxVal
	}
	return sum
}
