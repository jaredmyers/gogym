// Keywords: backtrack backtracking permutation
package main

// LEETCODE: 0046
// Medium
// time:O(), space: O(n!)
func permute(nums []int) [][]int {

	n := len(nums)
	result := [][]int{}

	var backtrack func(int)
	backtrack = func(first int) {
		if first == n {
			tmp := []int{}
			for _, val := range nums {
				tmp = append(tmp, val)
			}
			result = append(result, tmp)
		}

		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backtrack(0)
	return result
}
