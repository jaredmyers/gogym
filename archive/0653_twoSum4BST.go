package main

// LEETCODE: 0653

// time: O(n), space: O(n)
func findTarget(root *TreeNode, k int) bool {

	m := map[int]struct{}{}
	var find func(*TreeNode) bool

	find = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		_, ok := m[k-root.Val]
		if ok {
			return true
		}

		m[root.Val] = struct{}{}
		return find(root.Left) || find(root.Right)
	}

	return find(root)
}
