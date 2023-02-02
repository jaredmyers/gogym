package main

// LEEDCODE: 412

import "strconv"

// hashmap approach
// time: O(n), space: O(1)
func fizzBuzz2(n int) []string {

	return []string{}
}

// Naive approach
// time: O(n), space: O(1)
func fizzBuzz(n int) []string {

	answer := make([]string, n)
	for i := 1; i < n+1; i++ {

		if i%3 == 0 && i%5 == 0 {
			answer[i-1] = "FizzBuzz"
		}
		if i%3 == 0 {
			answer[i-1] = "Fizz"
		}
		if i%5 == 0 {
			answer[i-1] = "Buzz"
		}
		answer[i-1] = strconv.Itoa(i)
	}
	return answer
}
