package main

// LEETCODE: 226

// time: O(n), space: O(n)
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root

}
