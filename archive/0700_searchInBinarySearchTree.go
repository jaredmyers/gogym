package main

// LEETCODE: 700

// time: O(h), space: O(logn) worst: O(n)
func searchBST(root *TreeNode, val int) *TreeNode {

	node, _ := contains(root, val)
	return node
}

func contains(root *TreeNode, val int) (*TreeNode, bool) {
	if root.Val == val {
		return root, true
	}

	if val < root.Val {
		if root.Left == nil {
			return nil, false
		}
		node, found := contains(root.Left, val)
		if found {
			return node, found
		}
	}

	if val > root.Val {
		if root.Right == nil {
			return nil, false
		}
		node, found := contains(root.Right, val)
		if found {
			return node, found
		}
	}
	return nil, false
}
