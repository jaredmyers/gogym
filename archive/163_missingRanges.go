package main

// LEEDCODE: 163

import "strconv"

func findMissingRanges(nums []int, lower, upper int) []string {

	final := []string{}
	prev := lower - 1
	var curr int

	for i := 0; i <= len(nums); i++ {
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}

		if prev+1 <= curr-1 {
			p := strconv.Itoa(prev + 1)
			c := strconv.Itoa(curr - 1)
			if p == c {
				final = append(final, p)
			} else {
				final = append(final, (p + "->" + c))
			}
		}
		prev = curr
	}
	return final
}
