package main

// LEETCODE: 144

// time: O(n), space: O(1)
func preorderTraversal(root *TreeNode) []int {

	store := []int{}
	preOrder(root, &store)
	return store
}

func preOrder(node *TreeNode, store *[]int) {

	if node == nil {
		return
	}

	*store = append(*store, node.Val)

	if node.Left != nil {
		preOrder(node.Left, store)
	}
	if node.Right != nil {
		preOrder(node.Right, store)
	}

}
