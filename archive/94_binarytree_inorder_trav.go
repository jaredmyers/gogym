package main

// LEETCODE: 94

func inorderTaveral(root *TreeNode) []int {
	final := []int{}
	if root == nil {
		return final
	}

	trav(root, &final)
	return final
}

func trav(node *TreeNode, final *[]int) {

	if node == nil {
		return
	}
	trav(node.Left, final)
	*final = append(*final, node.Val)
	trav(node.Right, final)
}
