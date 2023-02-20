package main

// LEETCODE: 145

// time: O(n), space: O(n)
func postorderTraveral(root *TreeNode) []int {

	store := []int{}
	postOrder(root, &store)
	return store
}

func postOrder(root *TreeNode, store *[]int) {

	if root == nil {
		return
	}

	if root.Left != nil {
		postOrder(root.Left, store)
	}
	if root.Right != nil {
		postOrder(root.Right, store)
	}

	*store = append(*store, root.Val)
}
