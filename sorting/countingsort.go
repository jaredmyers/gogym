// Golang
// Counting Sort

// Time
// ====
// Best: O(n+k)
// Average: O(n+k)
// Worst: O(n+k)

// Space
// =====
// 0(max)

// Stable: Yes

package sorting

func CountingSort(arr []int, k int) []int {

	count := make([]int, k+1)

	// counts occurences of each
	// holds all elems form 0 to k-1
	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		count[arr[i]]++
	}

	//running sum
	// c[i] = num of occurences of values which are <= i
	for i := 1; i < len(count); i++ {
		count[i] = count[i-1] + count[i]
	}

	sorted := make([]int, len(arr))

	// linear pass through input arr backwards for stability
	// relevant indexes of count are values of arr[]
	// the index of count is the sorted placement of a[i]
	for i := len(arr) - 1; i >= 0; i-- {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	return sorted

}
