package main

// LEEDCODE: 2
// medium

// time: O(max(m, n)), space: O(max(m,n))
func addTwoNumbers(l1, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0
	var l1Val int
	var l2Val int

	for l1 != nil || l2 != nil || carry != 0 {

		if l1 != nil {
			l1Val = l1.Val
		} else {
			l1Val = 0
		}
		if l2 != nil {
			l2Val = l2.Val
		} else {
			l2Val = 0
		}

		columnSum := l1Val + l2Val + carry
		carry = columnSum / 10
		newNode := &ListNode{Val: (columnSum % 10)}
		curr.Next = newNode
		curr = newNode

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyHead.Next
}
