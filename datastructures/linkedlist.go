package datastructures

import "fmt"

type node struct {
		data string
		next *node
}

type LinkedList struct {
		head *node
		length int
}

func InitList() *LinkedList {
		return &LinkedList{}
}

func (s *LinkedList) Add(data string) {
		node := &node{
				data: data,
		}

		if s.head == nil {
				s.head = node
		}else{
				current := s.head
				for current.next != nil {
						current = current.next
				}
				current.next = node
		}
		s.length++
		return
}

func (s *LinkedList) Traverse() error {
		if s.head == nil {
				return fmt.Errorf("TraverseError: List Empty")
		}
		current := s.head
		for current != nil {
				fmt.Println(current.data)
				current = current.next
		}
		return nil
}
