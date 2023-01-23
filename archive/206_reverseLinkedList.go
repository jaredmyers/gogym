package main

// LEEDCODE: 206

// Time: O(n)
func reverseList2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	prev := &ListNode{Val: head.Val}
	curr := head.Next
	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}
	return prev

}

// dumb first try
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	tracker := []*ListNode{}
	curr := head
	for curr != nil {
		tracker = append(tracker, curr)
		curr = curr.Next
	}

	p1 := &ListNode{Val: tracker[len(tracker)-1].Val}
	tracker = tracker[:len(tracker)-1]
	next := p1

	for len(tracker != 0) {
		next.Next = &ListNode{Val: tracker[len(tracker)-1].Val}
		tracker = tracker[:len(tracker)-1]
		next = next.Next
	}
	return p1
}
