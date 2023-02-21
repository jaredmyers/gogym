package main

// LEETCODE: 0203

// time: O(n), space: O(1)
func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head

	prev, curr := dummy, dummy.Next
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}
