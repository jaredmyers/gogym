package main

// LEETCODE: 965

// time: O(N), space: O(h)
func isUnivalTree(root *TreeNode) bool {

	left := root.Left == nil || (root.Val == root.Left.Val && isUnivalTree(root.Left))
	right := root.Right == nil || (root.Val == root.Right.Val && isUnivalTree(root.Right))

	return left && right
}
