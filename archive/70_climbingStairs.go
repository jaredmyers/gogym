package main

// you are climbing stairs, it takes n steps to reach top
// each time you step you can either climb 1 or 2 steps.
// in how many distinct ways can you climb to the top?

//LEEDCODE: 70

func climbStairs(n int) int {

	if n < 1 || n > 45 {
		return 0
	}

	memo := make([]int, n+1)
	return climb(memo, 0, n)
}

func climb(memo []int, i, n int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if memo[i] > 0 {
		return memo[i]
	}

	memo[i] = climb(memo, i+1, n) + climb(memo, i+2, n)
	return memo[i]
}
