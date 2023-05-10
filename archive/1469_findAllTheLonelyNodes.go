package main

// LEEDCODE: 1469

// time: O(n), space: O(n)
func getLonelyNodes(root *TreeNodes) []int {

	if root == nil {
		return []int{}
	}

	tracker := []int{}
	dfs(root, &tracker)

	return tracker
}

func dfs(root *TreeNode, tracker *[]int) {

	if root == nil {
		return
	}

	if root.Left != nil && root.Right == nil {
		*tracker = append(*tracker, root.Left.Val)
	}
	if root.Left == nil && root.Right != nil {
		*tracker = append(*tracker, root.Right.Val)
	}

	dfs(root.Left, tracker)
	dfs(root.Right, tracker)
}
