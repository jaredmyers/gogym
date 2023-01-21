// GoLang
// Given array of ints and a target,
// return indices of two numbers such that
// they add up to the target

// Input: [2,7,11,15], target = 9
// output [0,1]

//LEETCODE: 1

package main

func twoSum(nums []int, target int) []int {

	m := map[int]int{}
	answer := []int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		value, ok := m[complement]
		if ok && value != i {
			answer = append(answer, i)
			answer = append(answer, value)
			return answer
		}
	}
	return answer
}
