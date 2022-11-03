// golang

package main

import (
	"fmt"
	//sort "github.com/jaredmyers/scratch/sorting"
	ds "github.com/jaredmyers/scratch/datastructures"
)

func main() {
		
		s := ds.InitStack()

		s.Push("Hello")
		s.Push("Testing")
		val, err := s.Peek()

		if err != nil {
				fmt.Println(err)
				
		}else{
				fmt.Println(val)
		}
}
