// Golang
// HeapSort

// Time
// ====
// Best: O(nlogn)
// Average: O(nlogn)
// Worst: O(nlogn)

// Space
// =====
// 0(n) in-place

// Stable: No

package sorting

import (
	"fmt"

	ds "github.com/jaredmyers/gogym/datastructures"
)

// using an anon func
func HeapSort(arr []int) {

	// first build max heap
	// input: arr, lastIdx, first node with children (last leaf node idx - 1)
	ds.BuildMaxHeap(arr)
	fmt.Println(arr)

	// linear pass backwards
	// sorting like insertion sort, but from back to front
	for i := len(arr); i > 1; i-- {

		// every pass takes in slice[n-1] of arr
		// segementing working elements (left side) from sorted elements (right side)
		// arr variable is our pointer to original array being sorted in-place
		func(slicer []int) {
			// swap root (largest with last)
			slicer[0], slicer[len(slicer)-1] = slicer[len(slicer)-1], slicer[0]
			//fmt.Println("sswapped:", slicer)

			// slice down to segement deleted elems from working elems
			slicer = slicer[:len(slicer)-1]

			// rebuild the max heap for the working elem portion (left side)
			ds.BuildMaxHeap(slicer)
			//fmt.Println("reheaped:", slicer)

		}(arr[:i])
	}
	fmt.Println(arr)
}

// without anon func
// explicitly recursive
func HeapSort2(arr []int) {

	// build maxheap
	ds.BuildMaxHeap(arr)
	fmt.Println(arr)
	// sort it
	sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) {

	// base case
	// not need to go to zero
	if len(arr) <= 1 {
		return
	}

	// swap root to back
	arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
	// slice down to segment
	arr = arr[:len(arr)-1]
	// rebuild max heap from right segmented side
	ds.BuildMaxHeap(arr)
	// go again
	sort(arr)

}
