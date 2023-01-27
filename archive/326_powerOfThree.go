package main

// LEETCODE: 326

// time: O(log3(n)), space: O(1)
func isPowerOfThree(n int) bool {

	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}
