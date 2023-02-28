// keywords: matrix, spiral
package main

// LEETCODE: 0059
// Medium

// time: O(mn), space: O(1)
func spiralOrder(matrix [][]int) []int {
	result := []int{}
	rows := len(matrix)
	cols := len(matrix[0])
	up, left, right, down := 0, 0, cols-1, rows-1

	for len(result) < rows*cols {
		//right
		for col := left; col <= right; col++ {
			result = append(result, matrix[up][col])
		}
		//down
		for row := up + 1; row <= down; row++ {
			result = append(result, matrix[row][right])
		}

		//left
		for col := right - 1; col >= left; col-- {
			result = append(result, matrix[down][col])
		}

		//up
		for row := down - 1; row > up; row-- {
			result = append(result, matrix[row][left])
		}

		left++
		right--
		up++
		down--
	}

	return result
}
