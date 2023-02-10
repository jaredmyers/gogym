package main

// LEETCODE: 0036
// Medium

// time: O(n^2), space: O(n^2) -- or O(1), O(1)
// since input is fixed
func isValidSudoku(board [][]byte) bool {

	n := 9

	rows := map[int]map[byte]bool{}
	cols := map[int]map[byte]bool{}
	boxes := map[int]map[byte]bool{}

	for r := 0; r < n; r++ {
		rows[r] = map[byte]bool{}
		cols[r] = map[byte]bool{}
		boxes[r] = map[byte]bool{}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {

			val := board[r][c]

			if string(val) == "." {
				continue
			}

			_, ok := rows[r][val]
			if ok {
				return false
			}
			rows[r][val] = true

			_, ok = cols[c][val]
			if ok {
				return false
			}
			cols[c][val] = true

			idx := (r/3)*3 + c/3
			_, ok = boxes[idx][val]
			if ok {
				return false
			}
			boxes[idx][val] = true

		}
	}
	return true
}
