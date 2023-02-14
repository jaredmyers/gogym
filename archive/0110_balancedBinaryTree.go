package main

import "math"

// LEEDCODE: 110

// time: O(n), space: O(n)

func isBalanced(root *TreeNode) bool {
	i, _ := isBalancedHelper(root)
	return i
}

func isBalancedHelper(root *Treenode) (bool, int) {

	if root == nil {
		return true, -1
	}

	leftIsBalanced, leftHeight := isBalancedHelper(root.Left)
	if !leftIsBalanced {
		return false, 0
	}
	rightIsBalanced, rightHeight := isBalancedHelper(root.Right)
	if !rightIsBalanced {
		return false, 0
	}

	var val bool
	if math.Abs(float64(leftHeight)-float64(rightHeight)) < 2 {
		val = true
	} else {
		val = false
	}

	return val, 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))

}
