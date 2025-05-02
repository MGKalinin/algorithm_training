// main_test.go
// go test -v запуск тестов
package main

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{
			input:  []int{1, 2, 3, 4},
			expect: []int{24, 12, 8, 6},
		},
		{
			input:  []int{-1, 1, 0, -3, 3},
			expect: []int{0, 0, 9, 0, 0},
		},
	}

	for i, test := range tests {
		result := productExceptSelf(test.input)
		if !slicesEqual(result, test.expect) {
			t.Errorf("Test %d failed:\nExpected: %v\nGot:      %v", i+1, test.expect, result)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
