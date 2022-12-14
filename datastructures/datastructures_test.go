package datastructures

import (
	"reflect"
	"testing"
)

var (
	arr []int = []int{15, 5, 20, 1, 17, 10, 30}
	cmp []int = []int{30, 17, 20, 1, 5, 10, 15}
)

func TestMaxHeapify(t *testing.T) {

	BuildMaxHeap(arr, len(arr)-1, (len(arr)/2)-1)
	if !reflect.DeepEqual(arr, cmp) {
		t.Errorf(" got %v", arr)
		t.Errorf(" want %v", cmp)
	}
}
