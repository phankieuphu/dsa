package test

import (
	"reflect"
	"testing"

	"github.com/phankieuphu/dsa.git/sort"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 1, 2}, []int{1, 2, 3, 5, 8}},
		{[]int{10, -1, 2, 5, 0}, []int{-1, 0, 2, 5, 10}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3, 3, 2, 1}, []int{1, 2, 3, 3}},
	}

	for _, test := range tests {
		result := sort.BubbleSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
