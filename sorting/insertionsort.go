// Golang
// Insertion Sort

// Time
// ====
// Best: O(n)
// Average: O(n^2)
// Worst: O(n^2)

// Space
// =====
// 0(n) in-place

// Stable: Yes

package sorting

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}
