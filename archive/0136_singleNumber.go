package main

// find single num in array of duplicate nums
// bit manipulation XOR
// LEEDCODE: 136

func singleNumber(nums []int) int {
	a := 0
	for _, val := range nums {
		a ^= val
	}
	return a
}

// in python
/*
def singleNumber(nums):
    a = 0
    for i in nums:
        print(a)
        a ^= i

    print(a)
*/
