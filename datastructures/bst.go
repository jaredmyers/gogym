// Golang
// BST

package datastructures

import "fmt"

type bstNode struct {
	data  int
	right *bstNode
	left  *bstNode
}

type BinarySearchTree struct {
	root *bstNode
}

func InitBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (b *BinarySearchTree) Insert(data int) error {

	if b.root != nil {
		err := b.root.insert(data)
		return err
	}

	b.root = &bstNode{data: data}
	return nil
}

func (n *bstNode) insert(data int) error {

	// if value less, go left
	if data < n.data {
		if n.left == nil {
			n.left = &bstNode{data: data}
		} else {
			n.left.insert(data)
		}
		return nil
	}

	// if value greater, go right
	// do not allow duplicate values
	if data > n.data {
		if n.right == nil {
			n.right = &bstNode{data: data}
		} else {
			n.right.insert(data)
		}
	}
	return nil
}

func (b *BinarySearchTree) Contains(data int) (bool, error) {

	if b.root != nil {
		contain := b.root.contains(data)
		return contain, nil
	}
	return false, fmt.Errorf("Contains Err: Tree Empty")
}

func (n *bstNode) contains(data int) bool {

	// when we find a match, return true
	if data == n.data {
		return true
	}

	// if data is less than current, but left is null
	// data not in tree, if left exists, dig
	if data < n.data {
		if n.left == nil {
			return false
		}
		// make sure bool survives recursion stack lol
		if n.left.contains(data) {
			return true
		}
	}
	// if data greater than current, but right is null
	// data not in tree, if right exists, dig
	if data > n.data {
		if n.right == nil {
			return false
		}
		// make sure bool survives recursion stack lol
		if n.right.contains(data) {
			return true
		}

	}

	return false
}

func (b *BinarySearchTree) TravLevelOrder() {

	// print nodes one layer at a time
	// BFS

	if b.root == nil {
		return
	}

	queue := []bstNode{}

	queue = append(queue, *b.root)
	b.root.travLevelOrder(b.root, queue)

}

func (n *bstNode) travLevelOrder(node *bstNode, q []bstNode) {

	// print + pop next off queue
	fmt.Print(q[0].data, "-")
	q = q[1:]

	// append to next layer unless leaf
	if node.left != nil {
		q = append(q, *node.left)
	}
	if node.right != nil {
		q = append(q, *node.right)
	}

	// if isEmpty even after next layer attempt
	// at end, return
	if len(q) == 0 {
		return
	}

	// go again
	node.travLevelOrder(&q[0], q)

}

func (b *BinarySearchTree) TravPostOrder() {

	// traverse left
	// then traverse right
	// then print values
	// (so print a value only after both its subtrees)
	// (have been explored)

	if b.root == nil {
		return
	}

	b.root.travPostOrder(b.root)
}

func (n *bstNode) travPostOrder(root *bstNode) {

	if root == nil {
		return
	}

	// postorder prints after recursive calls
	n.travPostOrder(root.left)
	n.travPostOrder(root.right)
	fmt.Print(root.data, "-")
}

func (b *BinarySearchTree) TravPreOrder() {

	// print value of current node,
	// then traverse left
	// then traverse right
	// (so dig left and print values on the way down)
	// (then pick up any rights on the way up)

	if b.root == nil {
		return
	}

	b.root.travPreOrder(b.root)
}

func (n *bstNode) travPreOrder(root *bstNode) {

	if root == nil {
		return
	}

	// preorder prints before recursive calls
	fmt.Print(root.data, "-")
	n.travPreOrder(root.left)
	n.travPreOrder(root.right)
}

func (b *BinarySearchTree) TravInOrder() {

	// traverse left,
	// print value,
	// then traverse right
	// (so dig all the way left, but print values on the way up)
	// (then pick up any rights on the way up)
	// or (print value after nodes left subtrees have been xplored)

	if b.root == nil {
		return
	}

	b.root.travInOrder(b.root)
}

func (n *bstNode) travInOrder(root *bstNode) {

	if root == nil {
		return
	}

	// inorder prints between recursive calls
	n.travInOrder(root.left)
	fmt.Print(root.data, "-")
	n.travInOrder(root.right)
}

// TO DO
// remove cases:
// 1. node to remove is a leaf node
// 2. node to remove has a right subtree, but no left subtree
// 3. node to remove has a left subtree, but no right subtree
// 4. node to remove has both left and right subtrees
func (b *BinarySearchTree) Remove(data int) {

}

func (n *bstNode) remove() {

}
