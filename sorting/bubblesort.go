package sorting

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
