package main

// LEEDCODE: 019
// Medium

// time: O(n), space: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	curr := head

	for i := 0; i < n; i++ {
		curr = curr.Next
	}

	if curr == nil {
		return head.Next
	}

	nodeBeforeRemoved := head
	for curr.Next != nil {
		curr = curr.Next
		nodeBeforeRemoved = nodeBeforeRemoved.Next
	}

	nodeBeforeRemoved.Next = nodeBeforeRemoved.Next.Next
	return head
}
