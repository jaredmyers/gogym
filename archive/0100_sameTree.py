# LEETCODE: 100

# time: O(n), space: O(n)
def isSameTree(p, q):

    if p == None and q == None:
        return True
    if p == None or q == None:
        return False

    if p.val != q.val:
        return False

    return isSameTree(p.right, q.right) and isSameTree(p.left, q.left)
