package main

// LEEDCODE: 1022

// time: O(n), space: O(h)
func sumRootToLeaf(root *TreeNode) int {

	rootToLeaf := 0
	var preorder func(*TreeNode, int)

	preorder = func(r *TreeNode, currNumber int) {

		if r != nil {
			currNumber = (currNumber << 1) | r.Val
			if r.Left == nil && r.Right == nil {
				rootToLeaf += currNumber
			}
			preorder(r.Left, currNumber)
			preorder(r.Right, currNumber)
		}
	}

	preorder(root, 0)
	return rootToLeaf
}
