// golang
// doubly linkedlist 

package datastructures

import "fmt"

type dllNode struct {
		data string
		next *dllNode
		prev *dllNode
}

type DoublyLinkedList struct {
		head *dllNode
		tail *dllNode
		length int
}

func InitDoublyLinkedList() *DoublyLinkedList {
		return &DoublyLinkedList{}
}

func (s *DoublyLinkedList) AddEnd(data string) {
		node := &dllNode{data: data,}

		if s.head == nil {
				s.head = node
				s.tail = node
		}else{
				node.prev = s.tail
				s.tail.next = node
				s.tail = node
		}
		s.length++
		return
}

func (s *DoublyLinkedList) AddFront(data string){
		
		node := &dllNode{data: data,}

		if s.head == nil {
				s.head = node
				s.tail = node
		}else{
				node.next = s.head
				s.head.prev = node
				s.head = node
		}
		s.length++
		return
}

func (s *DoublyLinkedList) RemoveEnd() error {
		
		if s.head == nil {
				return fmt.Errorf("RemoveError: List Empty")
		}

		current := s.tail
		current.prev.next = nil
		s.tail = current.prev

		return nil
}

func (s *DoublyLinkedList) RemoveFront() error {
		
		if s.head == nil {
				return fmt.Errorf("RemoveError: List Empty")
		}

		current := s.head
		current.next.prev = nil
		s.head = current.next

		return nil
}


func (s *DoublyLinkedList) TraverseForward() error {
		
		if s.head == nil{
				return fmt.Errorf("TraverseError: Empty List")
		}

		current := s.head

		for current != nil {
				fmt.Printf("%v->", current.data)
				current = current.next
		}
		fmt.Println("")
		return nil
}

func (s *DoublyLinkedList) TraverseReverse() error{
		if s.head == nil{
				return fmt.Errorf("TraverseError: List Empty")
		}

		current := s.tail

		for current != nil {
				fmt.Printf("%v->", current.data)
				current = current.prev
		}
		fmt.Println("")
		return nil
}

func (s *DoublyLinkedList) GetHead() string {
		return s.head.data
}
func (s *DoublyLinkedList) GetTail() string {
		return s.tail.data
}
