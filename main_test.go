package main

import (
	"fmt"
	"testing"
)

// функция тестирования prefixSum
func TestPrefixSum(t *testing.T) {
	tests := []struct {
		arr    []int
		n      int
		l      int
		r      int
		expect int
	}{

		{[]int{1, 2, 3, 4}, 4, 1, 1, 1},
		{[]int{1, 2, 3, 4}, 4, 1, 2, 3},
		{[]int{1, 2, 3, 4}, 4, 1, 3, 6},
		{[]int{1, 2, 3, 4}, 4, 1, 4, 10},
		{[]int{1, 2, 3, 4}, 4, 2, 2, 2},
		{[]int{1, 2, 3, 4}, 4, 2, 3, 5},
		{[]int{1, 2, 3, 4}, 4, 2, 4, 9},
		{[]int{1, 2, 3, 4}, 4, 3, 3, 3},
		{[]int{1, 2, 3, 4}, 4, 3, 4, 7},
		{[]int{1, 2, 3, 4}, 4, 4, 4, 4},
	}
	for i, test := range tests {
		result := prefixSum(test.arr, test.n, test.l, test.r)
		if result != test.expect {
			fmt.Printf("Тест %d провален: ожидалось %d, получено %d\n", i+1, test.expect, result)
		} else {
			fmt.Printf("Тест %d пройден\n", i+1)
		}
	}
}
