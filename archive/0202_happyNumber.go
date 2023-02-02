package main

// LEEDCODE: 202

import "fmt"

func isHappy(n int) bool {

	tracker := map[int]bool{}
	_, ok := tracker[n]
	for n != 1 && !ok {
		fmt.Println(n)
		tracker[n] = true
		n = func(n int) int {
			var total int
			var tmp int
			for n != 0 {
				tmp = n % 10
				n = n / 10
				total += tmp * tmp
			}
			return total

		}(n)

		_, ok = tracker[n]
	}

	return n == 1

}
