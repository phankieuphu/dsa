package sort

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := len(arr) / 2
	left := MergeSort(arr[:pivot])
	right := MergeSort(arr[pivot:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// options 1
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	// options 2
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result

}
