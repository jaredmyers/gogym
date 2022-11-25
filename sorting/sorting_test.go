// Golang
// sorting test

package sorting

import (
		"testing"
		"reflect"
)

var (
		arr []int = []int{12,234,10,1,67,94,19,4,100}
		cmp []int = []int{1,4,10,12,19,67,94,100,234}

)

func TestBubbleSort(t *testing.T) {
		
		BubbleSort(arr)
		if !reflect.DeepEqual(arr, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}

func TestSelectionSort(t *testing.T) {
		
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}

func TestInsertionSort(t *testing.T) {
		
		InsertionSort(arr)
		if !reflect.DeepEqual(arr, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}

func TestQuickSort(t *testing.T) {
		
		QuickSort(arr, 0, len(arr)-1)
		if !reflect.DeepEqual(arr, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}

func TestMergeSort(t *testing.T) {
		
		sorted := MergeSort(arr)
		if !reflect.DeepEqual(sorted, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}

func TestRadixSort(t *testing.T) {
		
		RadixSort(arr)
		if !reflect.DeepEqual(arr, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}

func TestCountingSort(t *testing.T) {
		
		CountingSort(arr, 234)
		if !reflect.DeepEqual(arr, cmp) {
				t.Errorf(" got %v", arr)
				t.Errorf(" want %v", cmp)
		}
}
