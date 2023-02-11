
# LEETCODE: 83

# time: O(n), space: O(1)
def deleteDuplicates(head):

    if head == None:
        return head

    p1 = head
    p2 = head.next

    while p2 != None:
        if p1.val == p2.val:
            p1.next = p2.next
            p2 = p2.next
            continue

        p1 = p1.next
        p2 = p2.next

    return head
