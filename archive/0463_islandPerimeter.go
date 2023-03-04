package main

// LEETCODE: 463

// time: O(mn), space: O(1)
func islandPerimeter(grid [][]int) int {

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 1 {
				continue
			}

			total += 4
			if i != 0 && grid[i-1][j] == 1 {
				total -= 2
			}
			if j != (len(grid[0])-1) && grid[i][j+1] == 1 {
				total -= 2
			}
		}
	}
	return total
}
