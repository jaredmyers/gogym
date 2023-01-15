package main

// LEETCODE: 141

// two pointer approach Time: O(n), Space O(1)
func hasCycle2(head *ListNode) bool {

	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// map approach  time: O(n), space: O(n)
func hasCycle(head *ListNode) bool {

	tracker := map[*ListNode]bool{}
	tracker[head] = true

	for head != nil {
		_, ok := tracker[head.Next]
		if ok {
			return true
		}

		tracker[head.Next] = true
		head = head.Next
	}
	return false
}
