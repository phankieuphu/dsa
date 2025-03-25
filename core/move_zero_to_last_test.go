package core

import (
	"reflect"
	"testing"
)

func TestMoveZeroToLast(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"No zeros", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"All zeros", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"Mixed numbers", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Leading zeros", []int{0, 0, 1, 2, 3}, []int{1, 2, 3, 0, 0}},
		{"Trailing zeros", []int{4, 5, 6, 0, 0}, []int{4, 5, 6, 0, 0}},
		{"Zeros in between", []int{1, 0, 2, 0, 3, 4}, []int{1, 2, 3, 4, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MoveZeroToLast(append([]int{}, tt.input...)) // Avoid modifying original input
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MoveZeroToLast(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
