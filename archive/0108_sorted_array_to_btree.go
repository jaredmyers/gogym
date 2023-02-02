package main

// given an sorted int array, convert it to a high-balanced btree
// LEETCODE: 108

func sortedArrayToBst(arr []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)

}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// always choose left middle as root
	p := (left + right) / 2

	// preorder traversal
	root := &TreeNode(nums[p])
	root.Left = helper(nums, left, p-1)
	root.Right = helper(nums, p+1, right)
	return root
}
