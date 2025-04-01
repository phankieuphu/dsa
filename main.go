package main

import (
	"fmt"

	"github.com/phankieuphu/dsa.git/sort"
)

func main() {
	input := []int{6, 3, 4, 1, 8, 5, 8, 9}
	result := sort.InsertionSort(input)
	fmt.Println("Sorted array:", result)
}
