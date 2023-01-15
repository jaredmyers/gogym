package main

// given root of Binary Tree, check if it is a mirror of itself

// LEETCODE: 101

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isMirror(left.Right, right.Left) && isMirror(left.Left, right.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
