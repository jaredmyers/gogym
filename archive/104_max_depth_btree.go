package main

// given root, return max depth
// LEETCODE: 104

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftHeight float64 = float64(maxDepth(root.Left))
	var rightHeight float64 = float64(maxDepth(root.Right))
	return int(math.Max(leftHeight, rightHeight)) + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
