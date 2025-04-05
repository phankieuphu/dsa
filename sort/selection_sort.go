package sort

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Assume the current position has the minimum
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum with the current position
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
