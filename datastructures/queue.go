// Golang
// Queue w/single linked list

package datastructures

import "fmt"

type qnode struct {
		data string
		next *qnode
}

type Queue struct {
		head *qnode
		tail *qnode
		length int
}

func InitQueue() *Queue {
		return &Queue{}
}

func (s *Queue) Size() int {
		return s.length
}

func (s *Queue) IsEmpty() bool{
		return s.Size() == 0
}

func (s *Queue) Peek() (string, error) {
		
		if s.IsEmpty(){
				return "", fmt.Errorf("PeekError: List Empty")
		}

		return s.head.data, nil
		
}

func (s *Queue) Poll() (string,error) {
		
		if s.IsEmpty() {
				return "", fmt.Errorf("PollError: List Empty")
		}

		data := s.head.data
		s.head = s.head.next
		s.length--

		return data, nil
}

func (s *Queue) Add(data string) error {
		
		node := &qnode{data: data,}

		if s.IsEmpty() {
				s.head = node
				s.tail = node
		} else {
				
				s.tail.next = node
				s.tail = node
		}
		s.length++

		return nil
}

func (s *Queue) Traverse() error {
		
		if s.IsEmpty() {
				return fmt.Errorf("TraverseError: List Empty")
		}

		current := s.head

		for current != nil {
				
				fmt.Printf("%v <- ", current.data)
				current = current.next
		}
		fmt.Println("")
		
		return nil
}
