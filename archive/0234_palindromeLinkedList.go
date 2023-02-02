package main

// LEEDCODE: 234

// split list in half
// reverse last last half
// compare
// restore list
// Time: O(n), Space: O(1)

func isPalindrom2(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfStart
	result := true

	for result && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// restore list
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {

	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {

	prev := &ListNode{Val: head.Val}
	curr := head.Next

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

// first try
func isPalindrome(head *ListNode) bool {

	tracker := []int{head.Val}
	prev := &ListNode{Val: head.Val}
	curr := head.Next

	for curr != nil {
		tracker = append(tracker, curr.Val)
		tmpNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNext
	}

	p1 := prev
	i := 0

	for p1 != nil {
		if p1.Val != tracker[i] {
			return false
		}
		p1 = p1.Next
		i++
	}
	return true
}
