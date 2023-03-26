package main

// LEETCODE: 637

// time: O(n), space: O(m) m = max(level)
func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	q := []*TreeNode{}
	q = append(q, root)

	for len(q) > 0 {
		sum, count := float64(0), float64(0)
		temp := []*TreeNode{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			sum += float64(node.Val)
			count++
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		q = temp
		result = append(result, (sum / count))
	}
	return result
}
