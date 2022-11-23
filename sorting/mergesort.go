// Golang
// Merge Sort

// Time
// ====
// Best: O(nlogn)
// Average: O(nlogn)
// Worst: O(nlogn)

// Space
// =====
// 0(n) not in place

// Stable: Yes

package sorting

func MergeSort(items []int) []int {

	if len(items) < 2 {
		return items
	} //base case

	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])

	return merge(first, second)
}

func merge(a, b []int) []int {

	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}