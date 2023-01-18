package main

// LEEDCODE: 171

import "math"

// first go
func titleToNumber(columnTitle string) int {

	letters := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}

	final := 0
	var i float64 = 0
	var base float64 = 26
	for j := len(columnTitle) - 1; j >= 0; j-- {
		final += (letters[string(columnTitle[j])] * int(math.Pow(base, i)))
		i++
	}

	return final
}
