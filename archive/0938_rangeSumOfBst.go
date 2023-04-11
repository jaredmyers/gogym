package main

// LEEDCODE: 938

// time: O(n), space: O(n)
func rangeSumBST(root *TreeNode, low, high int) int {

	if root == nil {
		return 0
	}

	return dfs(root, low, high)
}

func dfs(root *TreeNode, low, high int) int {

	if root == nil {
		return 0
	}

	tracker := 0

	if root.Val >= low && root.Val <= high {
		tracker += root.Val
	}

	tracker += dfs(root.Left, low, high)
	tracker += dfs(root.Right, low, high)

	return tracker

}
