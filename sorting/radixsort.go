// Golang
// LSD Radix Sort
// w/ 2d array

// Time (figure out later)
// ====
// Best: O()
// Average: O()
// Worst: O()

// Space
// =====
// 0()

// Stable: Yes

package sorting

func RadixSort(arr []int) {

    // linear scan
	// find max num
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	divisor := 1
	for max/divisor > 0 {

		buckets := [10][]int{}

		// fill buckets
		for i := 0; i < len(arr); i++ {

			bucketIdx := (arr[i] / divisor) % 10
			buckets[bucketIdx] = append(buckets[bucketIdx], arr[i])
		}

		//empty buckets
		p := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < len(buckets[i]); j++ {
				arr[p] = buckets[i][j]
				p++
			}

		}
		//move to next significant digit/divisor
		divisor *= 10
	}
}
