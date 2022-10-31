//go practice
// chap4 of C++ gaddis problems

package main

import "fmt"

type node struct {
		data string
		next *node
		prev *node
}

type linkedList struct {
		head *node
		tail *node
		length int
}

func initList() *linkedList {
		return &linkedList{}
}

func (s *linkedList) AddEnd(data string) {
		node := &node{data: data,}

		if s.head == nil {
				s.head = node
				s.tail = node
		}else{
				current := s.head

				for current.next != nil {
						current = current.next
				}
				node.prev = current
				current.next = node
				s.tail = node
		}
		s.length++
		return
}

func (s *linkedList) AddFront(data string){
		
		node := &node{data: data,}

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


func (s *linkedList) TraverseForward() error {
		
		if s.head == nil{
				return fmt.Errorf("TraverseError: Empty List")
		}

		current := s.head

		for current != nil {
				fmt.Println(current.data)
				current = current.next
		}
		return nil
}

func (s *linkedList) TraverseReverse() error{
		if s.head == nil{
				return fmt.Errorf("TraverseError: List Empty")
		}

		current := s.tail

		for current != nil {
				fmt.Println(current.data)
				current = current.prev
		}
		return nil
}

func (s *linkedList) GetHead() string {
		return s.head.data
}
func (s *linkedList) GetTail() string {
		return s.tail.data
}

func main(){

		doublyList := initList()
		doublyList.AddEnd("Hello")
		doublyList.AddEnd("this")
		doublyList.AddEnd("is")
		doublyList.AddEnd("a")
		doublyList.AddEnd("test.")

		doublyList.TraverseForward()
		fmt.Println("---")
		doublyList.TraverseReverse()
		
}

