// golang
// doubly linkedlist 

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
				node.prev = s.tail
				s.tail.next = node
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

func (s *linkedList) RemoveEnd() error {
		
		if s.head == nil {
				return fmt.Errorf("RemoveError: List Empty")
		}

		current := s.tail
		current.prev.next = nil
		s.tail = current.prev

		return nil
}

func (s *linkedList) RemoveFront() error {
		
		if s.head == nil {
				return fmt.Errorf("RemoveError: List Empty")
		}

		current := s.head
		current.next.prev = nil
		s.head = current.next

		return nil
}


func (s *linkedList) TraverseForward() error {
		
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

func (s *linkedList) TraverseReverse() error{
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
		doublyList.AddEnd("test")

		doublyList.TraverseForward()
		doublyList.TraverseReverse()
		fmt.Println("Head: ",doublyList.GetHead())
		fmt.Println("Tail: ",doublyList.GetTail())

		doublyList.RemoveEnd()
		doublyList.TraverseForward()
		fmt.Println(doublyList.GetTail())

		doublyList.RemoveFront()
		doublyList.TraverseForward()
		fmt.Println(doublyList.GetHead())
		
}

