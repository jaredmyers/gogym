package main

// LEETCODE: 0590

// time: O(n), space: O(n)
func postOrder(root *Node) []int {

	store := []int{}
	q := []*Node{}
	traverse(root, &store, q)
	return store
}

func traverse(root *Node, store *[]int, q []*Node) {

	if root == nil {
		return
	}

	for _, val := range root.Children {
		q = append(q, val)
	}

	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		newQ := []*Node{}
		traverse(node, store, newQ)
	}

	*store = append(*store, root.Val)
}
