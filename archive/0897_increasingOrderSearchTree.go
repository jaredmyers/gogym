package main

// LEETCODE: 897

// time: O(n), space: O(h) height of tree
func increasingBST(root *TreeNode) *TreeNode {

	newRoot := &TreeNode{}
	//inorder dfs
	inorderDFS(root, newRoot)
	return newRoot.Right
}

func inorderDFS(node, curr *TreeNode) *TreeNode {

	if node == nil {
		return curr
	}

	curr = inorderDFS(node.Left, curr)
	node.Left = nil
	curr.Right = node
	curr = node
	curr = inorderDFS(node.Right, curr)
	return curr
}
