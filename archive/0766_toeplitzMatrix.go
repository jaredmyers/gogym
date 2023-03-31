package main

// LEETCODE: 0776

// time: O(mn), space: O(1)
func isToeplitzMatrix(matrix [][]int) bool {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i > 0 && j > 0 && matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}
	return true
}

// first try, exceeded time limit
func isToeplitzMatrix2(matrix [][]int) bool {

	var scan func([][]int, int, int) bool
	scan = func(matrix [][]int, x int, y int) bool {

		n := len(matrix)
		m := len(matrix[0])

		if x >= n || y >= m {
			return true
		}

		i := x
		j := y
		cmp := matrix[x][y]
		for ; i < n; i, j = i+1, j+1 {
			if j == m {
				break
			}
			if cmp != matrix[i][j] {
				return false
			}
		}
		return scan(matrix, x+1, y) && scan(matrix, x, y+1)
	}

	return scan(matrix, 0, 0)

}
