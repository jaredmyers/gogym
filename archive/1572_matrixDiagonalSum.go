package main

// LEETCODE: 1572

// time: O(n), space: O(1)
func diagonalSum(mat [][]int) int {

	middle := 0
	n := len(mat)
	if n%2 != 0 {
		middle = mat[n/2][n/2]
	}

	sum := 0
	j := n - 1
	for i := 0; i < n; i++ {
		sum += mat[i][i] + mat[i][j]
		j--
	}
	sum = sum - middle
	return sum
}
