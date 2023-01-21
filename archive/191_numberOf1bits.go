package main

// LEETCODE 191

func hammingWeight(num uint32) int {

	base := 32
	count := 0
	temp := uint32(0)

	for i := 0; i < base; i++ {
		temp = (num & 1)
		num = num >> 1
		if temp == 1 {
			count += 1
		}
	}
	return count
}
