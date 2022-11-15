// Golang
// Selection Sort

// Time
// ====
// Best: O(n^2)
// Average: O(n^2)
// Worst: O(n^2)

// Space
// =====
// 0(n) in-place

// Stable: Yes

package sorting

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
