package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		// TODO use n,k
		in   []int
		min  int
		list []int
	}{
		{[]int{7, 2}, 2, []int{1, 6}},
		{[]int{5, 1}, 2, []int{1, 4}},
		{[]int{6, 1}, 2, []int{2, 5}},
		{[]int{10, 0}, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{10, 1}, 4, []int{1, 4, 7, 10}},
		{[]int{100, 2}, 20, []int{3, 8, 13, 18, 23, 28, 33, 38, 43, 48, 53, 58, 63, 68, 73, 78, 83, 88, 93, 98}},
		{[]int{10, 10}, 1, []int{10}},
	}

	for i, test := range tests {
		min, list := solve(test.in[0], test.in[1])
		fmt.Println("list result", list)
		if min != test.min {
			t.Errorf("case %d expected min %d, got %d", i, test.min, min)
		}
		for j := range test.list {
			if list[j] != test.list[j] {
				t.Errorf("case %d expected list[%d] %d, got %d", i, j, test.list[j], list[j])
			}
		}
	}
}
