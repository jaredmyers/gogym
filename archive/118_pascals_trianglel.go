package main

// LEETCODE: 118

// rather than working element at a time, this solution
// works row at a time, and keeps track of two rows (curr and prev) on each iteration
// it caluculates each elem in a row based on the previous rows elems
// the algo also only focuses on the inner elems, and assigns 1 to each rows outer elems
// before and after the algorithm
// then just append the row array to the final array when completed the row

func generatePascalTriangle2(numRows int) [][]int {
	final := [][]int{}

	// base case, first row always has 1
	final = append(final, []int{})
	final[0] = append(final[0], 1)

	for rowNum := 1; rowNum < numRows; rowNum++ {

		prevRow := final[rowNum-1]
		row := []int{}

		// first row elem is always 1
		row = append(row, 1)

		// each triangle elem (than than first and last of each row)
		// is equal to the sum of the elems above-and-to-the-left and
		// above-and-to-the-right
		for i := 1; i < rowNum; i++ {
			row = append(row, prevRow[i-1]+prevRow[i])
		}
		// the last row elem is always 1
		row = append(row, 1)
		final = append(final, row)
	}
	return final
}

// my over worked solution. can't complete.
func generatePascalTriangle(numRows int) [][]int {

	final := [][]int{}
	innerIdx := 5
	innerIdxInterval := 1
	innerIdxTracker := 0
	for i := 0; i < numRows; i++ {
		final = append(final, []int{})
		j := i * 1
		for p := 0; p <= j; p++ {
			innerIdxTracker += 1
			if innerIdxTracker == innerIdx {
				for x := 0; x < innerIdxInterval; x++ {
					final[i] = append(final[i], final[i-1][p-1]+final[i-1][p])
				}
				innerIdx += 3
				innerIdxInterval += 1
			} else {
				final[i] = append(final[i], 1)
			}
		}
	}

	return final

}
