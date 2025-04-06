package main

import (
	"fmt"

	"github.com/phankieuphu/dsa.git/sliding_window"
)

func main() {
	//input := []int{6, 3, 4, 1, 8, 5, 8, 9}
	inputString := "abcdab"
	result := sliding_window.SlidingWindow(inputString)
	fmt.Println("Sorted array:", result)
}
