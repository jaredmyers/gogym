package main

// intersection of a linkedlist
// LEEDCODE: 161

// two pointer approach, time: O(N + M), space: O(1)
func getInsersectNode(headA, headB *ListNode) *ListNode {

	pA := headA
	pB := headB

	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}

// map approach  time: O(M) + O(N)
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	tracker := map[*ListNode]bool{}

	for headA != nil {
		tracker[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		_, ok := tracker[headB]
		if ok {
			return headB
		}
	}
	return nil
}
