// go test -v запуск; перейти в директорию с кодом))
package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "basic case",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "no zeros",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "zeros already at end",
			input:    []int{1, 2, 0, 0},
			expected: []int{1, 2, 0, 0},
		},
		{
			name:     "zeros at beginning",
			input:    []int{0, 0, 1, 2},
			expected: []int{1, 2, 0, 0},
		},
		{
			name:     "single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "single non-zero",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "random zeros",
			input:    []int{4, 0, 2, 0, 5},
			expected: []int{4, 2, 5, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем копию входного массива
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			// Вызываем тестируемую функцию
			moveZeroes(inputCopy)

			// Проверяем результат
			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("For input %v, expected %v but got %v",
					tt.input, tt.expected, inputCopy)
			}
		})
	}
}
