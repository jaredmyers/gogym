package main

// LEEDCODE: 872

// time: O(n), space: O(Max(leaves))
func leafSimilar(root1, root2 *TreeNode) bool {

	r1Queue := []int{}
	r2Queue := []int{}
	dfs(root1, &r1Queue)
	dfs(root2, &r2Queue)

	if len(r1Queue) != len(r2Queue) {
		return false
	}

	for i, val := range r1Queue {
		if r2Queue[i] != val {
			return false
		}
	}

	return true
}

func dfs(root *TreeNode, q *[]int) {

	if root.Left == nil && root.Right == nil {
		*q = append(*q, root.Val)
		return
	}

	if root.Left != nil {
		dfs(root.Left, q)
	}
	if root.Right != nil {
		dfs(root.Right, q)
	}
}
