// Golang
// Quick Sort

// Time
// ====
// Best: O(nlogn)
// Average: O(nlogn)
// Worst: O(n^2)

// Space
// =====
// 0(logn) in-place

// Stable: No

package sorting

func QuickSort(arr []int, start int, end int) {

	if end <= start {
		return
	}

	pivotIdx := partition(arr, start, end)
	QuickSort(arr, start, pivotIdx-1)
	QuickSort(arr, pivotIdx+1, end)
}

func partition(arr []int, start int, end int) int {

	pivot := arr[end]
	i := start - 1

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	return i

}
