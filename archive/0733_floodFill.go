package main

// LEETCODE: 0773

// time: O(n), space: O(n)
func floodFill(image [][]int, sr, sc, color int) [][]int {

	rows, cols := len(image), len(image[0])
	oldColor := image[sr][sc]
	if oldColor == color {
		return image
	}

	var dfs func(int, int)
	dfs = func(r int, c int) {
		if image[r][c] == oldColor {
			image[r][c] = color
		}
		if r >= 1 {
			dfs(r-1, c)
		}
		if r+1 < rows {
			dfs(r+1, c)
		}
		if c >= 1 {
			dfs(r, c-1)
		}
		if c+1 < cols {
			dfs(r, c+1)
		}
	}

	dfs(sr, sc)
	return image

}
