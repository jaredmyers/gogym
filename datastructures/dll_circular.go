//circular doubly linkedlist
//Golang

package datastructures

import "fmt"

type cllnode struct {
	data string
	next *cllnode
	prev *cllnode
}

type CircularList struct {
	head   *cllnode
	tail   *cllnode
	length int
}

func InitCircularList() *CircularList {
	return &CircularList{}
}

func (s *CircularList) AddEnd(data string) {

	node := &cllnode{data: data}

	if s.head == nil {
		s.head = node
		s.tail = node

		s.head.prev = s.tail
		s.tail.next = s.head
	} else {
		node.prev = s.tail
		s.tail.next = node
		s.tail = node

		// make circular
		s.tail.next = s.head
		s.head.prev = s.tail

	}
	s.length++
	return
}

func (s *CircularList) AddFront(data string) {
	node := &cllnode{data: data}

	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		node.next = s.head
		s.head.prev = node
		s.head = node

		//make circular
		s.head.prev = s.tail
		s.tail.next = s.head
	}
	s.length++
	return
}

func (s *CircularList) TraverseForward() error {
	if s.head == nil {
		return fmt.Errorf("TraverseError: List Empty")
	}

	current := s.head

	for {
		fmt.Printf("%v->", current.data)
		current = current.next
		if current == s.head {
			break
		}
	}
	return nil
}

func (s *CircularList) TraverseReverse() error {
	if s.head == nil {
		return fmt.Errorf("TraverseError: List Empty")
	}

	current := s.tail

	for {
		fmt.Printf("%v->", current.data)
		current = current.prev
		if current == s.tail {
			break
		}
	}
	return nil
}

func (s *CircularList) GetHead() string {
	return s.head.data
}

func (s *CircularList) GetTail() string {
	return s.tail.data
}
