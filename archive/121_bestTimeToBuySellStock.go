package main

import (
	"math"
)

// given array of prices where i is the i-th day,
// maxmize profit by buying stock as low as possible
// then selling as high as possible

// LEETCODE: 121

func maxProfit(prices []int) int {

	maxProfit := 0
	minPrice := math.MaxInt64

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
