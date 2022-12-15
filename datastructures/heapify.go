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

func BuildMaxHeap(arr []int) {

	// i = first node with children = (last leafIdx-1)
	i := (len(arr) / 2) - 1

	// run linear pass backwards on only parent (non-leaf) nodes
	for ; i >= 0; i-- {
		maxHeapify(arr, len(arr)-1, i)
	}

}

func maxHeapify(arr []int, last, i int) {

	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// check if potential child is inbounds/exists
	// check if child is larger than potential parent, tag largest if true
	if left <= last && arr[left] > arr[largest] {
		largest = left
	}

	// check if potential child is inbounds/exists
	// check if child is larger than potential parent, tag largest if true
	if right <= last && arr[right] > arr[largest] {
		largest = right
	}

	// after final determination of largest make parent/child swap
	// if swap true, recursively dig into tree to reheapify subtree
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		maxHeapify(arr, last, largest)
	}
}

func BuildMinHeap(arr []int) {

}

func minHeapify(arr []int) {

}
