// Golang
// file fun
// r/w json
// marshal/unmarshal testing

package main

import (
	"fmt"
)

type digits interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type bstNode[T digits] struct {
	data  T
	right *bstNode[T]
	left  *bstNode[T]
}

type BinarySearchTree[T digits] struct {
	root *bstNode[T]
}

func InitBST[T digits]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (b *BinarySearchTree[T]) Insert(data T) error {

	if b.root != nil {
		err := b.root.insert(data)
		return err
	}

	b.root = &bstNode[T]{data: data}
	return nil
}

func (n *bstNode[T]) insert(data T) error {

	if data < n.data {
		if n.left == nil {
			n.left = &bstNode[T]{data: data}
		} else {
			n.left.insert(data)
		}
		return nil
	}

	if data > n.data {
		if n.right == nil {
			n.right = &bstNode[T]{data: data}
		} else {
			n.right.insert(data)
		}
	}
	return nil
}

func (b *BinarySearchTree[T]) TravPostOrder() {

	if b.root == nil {
		return
	}

	b.root.travPostOrder(b.root)

}

func (n *bstNode[T]) travPostOrder(root *bstNode[T]) {

	if root == nil {
		return
	}

	n.travPostOrder(root.left)
	n.travPostOrder(root.right)
	fmt.Print(root.data, "-")

}

func main() {

}
