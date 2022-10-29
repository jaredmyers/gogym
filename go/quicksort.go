// quicksort

package main

func quickSort(arr []int, start int, end int) {

	if end <= start {
		return
	}

	pivotIdx := partition(arr, start, end)
	quickSort(arr, start, pivotIdx-1)
	quickSort(arr, pivotIdx+1, end)
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
