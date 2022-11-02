package main

import "fmt"

type node struct {
		data string
		next *node
}

type linkedList struct {
		head *node
		length int
}

func initList() *linkedList {
		return &linkedList{}
}

func (s *linkedList) Add(data string) {
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

func (s *linkedList) Traverse() error {
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
