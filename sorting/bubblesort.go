// Golang
// Bubble Sort

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

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
