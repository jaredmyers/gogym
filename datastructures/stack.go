// golang
// simple stack	

package datastructures

import "fmt"

type item struct {
		data string
		next *item
}

type Stack struct {
		top *item
		size int
}

func InitStack() *Stack {
		return &Stack{}
}

func (s *Stack) Size() int{
		return s.size
}

func (s *Stack) IsEmpty() bool{
		return s.size == 0
}

func (s *Stack) Push(data string) {
		
		s.top = &item{
				data: data,
				next: s.top,
		}
		s.size++
}

func (s *Stack) Pop() error {
	if s.size > 0 {
		s.top = s.top.next
		s.size--
	}else{
		return fmt.Errorf("PopError: Stack Empty")
	}
	return nil

}

func (s *Stack) Peek() (string, error) {

		if s.size > 0 {
				return s.top.data, nil
		}else{
				return "", fmt.Errorf("PeekError: Stack Empty")
		}

}
