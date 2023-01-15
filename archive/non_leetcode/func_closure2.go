// Golang
// function closure
// is a function which returns an anon func
// where the returned func encapsulates over
// an outer scope var from its parent function
// it "closes over the variable"
// effectively allowing state storage

package main

import "fmt"

func notmain() {

	f := foo()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func foo() func() int {

	x := 0
	return func() int {
		x++
		return x
	}
}
