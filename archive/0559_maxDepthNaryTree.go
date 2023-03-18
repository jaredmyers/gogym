package main

// LEETCODE: 0559

// time: O(n), space: O(n)
func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}

	if len(root.Children) == 0 {
		return 1
	}

	heights := []int{}
	for _, val := range root.Children {
		heights = append(heights, maxDepth(val))
	}

	max := heights[0]
	for _, val := range heights {
		if max < val {
			max = val
		}
	}
	return max + 1
}
