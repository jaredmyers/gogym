// Golang
// Binary MinHeap w/array

package datastructures

import "fmt"

type BinaryHeap struct {
	heap []int
}

func InitBinaryHeap() *BinaryHeap {
	return &BinaryHeap{}
}

func (b *BinaryHeap) Size() int {
	return len(b.heap)
}

func (b *BinaryHeap) Add(data int) {

	// add to end of tree
	b.heap = append(b.heap, data)
	// reorder tree
	b.bubbleUp(b.Size() - 1)
}

func (b *BinaryHeap) bubbleUp(idx int) {

	for idx > 0 {

		// get parent idx of newly inserted elem
		parentIdx := (idx - 1) / 2

		// if parent is less, min heap satisifed, break out
		if b.heap[parentIdx] < b.heap[idx] {
			return
		}

		//if parent is greater, swap & repeat
		b.heap[parentIdx], b.heap[idx] = b.heap[idx], b.heap[parentIdx]
		idx = parentIdx
	}
}

func (b *BinaryHeap) Poll() (int, error) {

	if b.Size() <= 0 {
		return 1, fmt.Errorf("PollError: Heap Empty")
	}
	polled := b.heap[0]
	size := b.Size()

	// swap with end of tree
	if size > 1 {
		b.heap[0] = b.heap[len(b.heap)-1]
	}

	// remove duplicate from end of tree
	b.heap = b.heap[:len(b.heap)-1]

	// reorder tree
	b.bubbleDown(0)

	return polled, nil
}

func (b *BinaryHeap) Remove(data int) error {

	if b.Size() <= 0 {
		return fmt.Errorf("RemoveError: Heap Empty")
	}

	// linear scan to find data elem
	// if size != zero and number greater than root
	if b.heap[0] <= data && b.Size() >= 1 {
		i := 0
		for i <= len(b.heap) {
			// return if number not detected
			if i == len(b.heap) {
				return nil
			}
			// break if index is found
			if b.heap[i] == data {
				break
			}
			i++
		}

		//swap with last elem
		b.heap[i], b.heap[len(b.heap)-1] = b.heap[len(b.heap)-1], b.heap[i]

		// remove elem from end of tree
		b.heap = b.heap[:len(b.heap)-1]

		// reorder tree (bubble swapped elem)
		// bubble up if parent is greater, else try to bubble down
		parentIdx := (i - 1) / 2
		if parentIdx >= 1 && b.heap[parentIdx] > b.heap[i] {
			b.bubbleUp(i)
		} else {
			b.bubbleDown(i)
		}

	}

	return nil
}

func (b *BinaryHeap) bubbleDown(idx int) {
	// if child exists then bubble down where applicable
	for (2*idx + 1) < len(b.heap) {

		// grab index of child which is lesser in value
		minChildIdx := b.minChildIdx(idx)

		// if min child is larger than the current parent
		// min heap satisfied, break out
		// if not, swap & repeat
		if b.heap[minChildIdx] >= b.heap[idx] {
			return
		}

		//swap
		b.heap[minChildIdx], b.heap[idx] = b.heap[idx], b.heap[minChildIdx]

	}
}

func (b *BinaryHeap) minChildIdx(idx int) int {

	// if right child doesn't exist, return left child index
	if (2*idx + 2) >= len(b.heap) {
		return 2*idx + 1
	}

	// since right child exists, return its index if smaller than left child
	if b.heap[2*idx+2] < b.heap[2*idx+1] {
		return 2*idx + 2
	}

	// if right child exists but is larger, return left child
	return 2*idx + 1
}

func (b *BinaryHeap) TmpTraverse() error {

	if b.Size() == 0 {
		return fmt.Errorf("TraverseError: Heap Empty")
	}

	for idx, val := range b.heap {
		fmt.Printf("%d:%d ", idx, val)
	}
	fmt.Println()
	return nil
}
