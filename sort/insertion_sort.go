package sort

func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			if arr[j] > key {
				arr[j+1] = arr[j]
				j = j - 1
			}
		}
		arr[j+1] = key
	}
	return arr
}
