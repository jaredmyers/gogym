package main

// LEETCODE: 1672

// time: O(MN), space: O(1)
func maximumWealth(accounts [][]int) int {

	max := 0
	for i := 0; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
