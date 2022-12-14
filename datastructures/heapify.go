//Golang
// Max & Min Heapify

// heapify = process of rearranging heap by comparing each
// largest with its children recursively.

// notes to future self
// heapsize = N
// range of leaves = (N/2) to (N-1)

// Heap Property
// maxheap - root node greater then all left/right subtrees
// recursively true for all subtrees
// minheap - root node less then all left/right subtrees
// recursively true for all subtrees

package datastructures

func BuildMaxHeap(arr []int, last, i int) {

	for ; i >= 0; i-- {
		maxHeapify(arr, len(arr)-1, i)
	}

}

func maxHeapify(arr []int, last, i int) {

	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left <= last && arr[left] > arr[largest] {
		largest = left
	}

	if right <= last && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		maxHeapify(arr, last, largest)
	}
}

func BuildMinHeap(arr []int) {

}

func MinHeapify(arr []int) {

}
