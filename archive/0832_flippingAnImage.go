package main

// LEETCODE: 832

// time: O(n), space: O(1)
func flipAndInvertImage(image [][]int) [][]int {

	n := len(image)
	for i := 0; i < n; i++ {
		p := n - 1
		for j := 0; j < n; j++ {
			if j < n/2 {
				image[i][j], image[i][p] = image[i][p], image[i][j]
				p--
			}
			if image[i][j] == 1 {
				image[i][j] = 0
			} else {
				image[i][j] = 1
			}
		}
	}
	return image
}
