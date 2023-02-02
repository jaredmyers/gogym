package main

// given non-neg int x, return the sqroot rounded down to nearest integer.
// dont use any language built ins.
// LEEDCODE: 69

func mySqrt(x int) int {

	if x < 2 {
		return x
	}

	var pivot int
	left := 2
	right := x / 2

	for left <= right {

		pivot = left + (left-right)/2
		num := pivot * pivot
		if num > x {
			right = pivot - 1
		} else if num < x {
			left = pivot + 1
		} else {
			return pivot
		}
	}
	return right
}
