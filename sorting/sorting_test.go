// Golang
// sorting test

package sorting

import (
	"reflect"
	"testing"
)

var (
	test   []int = []int{12, 234, 10, 1, 67, 94, 19, 4, 100}
	cptest []int = []int{12, 234, 10, 1, 67, 94, 19, 4, 100}
	cmp    []int = []int{1, 4, 10, 12, 19, 67, 94, 100, 234}

	h         []int = []int{15, 5, 20, 1, 17, 10, 30}
	cph       []int = []int{15, 5, 20, 1, 17, 10, 30}
	heapified []int = []int{30, 17, 20, 1, 5, 10, 15}
	hsorted   []int = []int{1, 5, 10, 15, 17, 20, 30}
)

func TestBubbleSort(t *testing.T) {

	BubbleSort(test)
	if !reflect.DeepEqual(test, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestSelectionSort(t *testing.T) {

	test = cptest
	SelectionSort(test)
	if !reflect.DeepEqual(test, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestInsertionSort(t *testing.T) {

	test = cptest
	InsertionSort(test)
	if !reflect.DeepEqual(test, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestQuickSort(t *testing.T) {

	test = cptest
	QuickSort(test, 0, len(test)-1)
	if !reflect.DeepEqual(test, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestMergeSort(t *testing.T) {

	test = cptest
	sorted := MergeSort(test)
	if !reflect.DeepEqual(sorted, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestRadixSort(t *testing.T) {

	test = cptest
	RadixSort(test)
	if !reflect.DeepEqual(test, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestCountingSort(t *testing.T) {

	test = cptest
	CountingSort(test, 234)
	if !reflect.DeepEqual(test, cmp) {
		t.Errorf(" got %v", test)
		t.Errorf(" want %v", cmp)
	}
}

func TestHeapSort(t *testing.T) {

	HeapSort(h)
	if !reflect.DeepEqual(h, hsorted) {
		t.Errorf(" got %v", h)
		t.Errorf(" want %v", hsorted)
	}

	h = cph
	HeapSort2(h)
	if !reflect.DeepEqual(h, hsorted) {
		t.Errorf(" got %v", h)
		t.Errorf(" want %v", hsorted)
	}

}
