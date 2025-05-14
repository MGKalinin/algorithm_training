package main

import (
	"fmt"
	"github.com/stretchr/testify/assert" // Можно использовать assert для удобства
	"testing"
)

// TestSliceDig проверяет базовые случаи
func TestSliceDig(t *testing.T) {
	// Тест 1: N = 5
	t.Run("N=5", func(t *testing.T) {
		result := sliceDig(5)
		expected := []int{0, 1, 2, 3, 4}
		assert.Equal(t, expected, result, "Слайс должен содержать числа 0-4")
		assert.Equal(t, 5, len(result), "Длина слайса должна быть 5")
	})

	// Тест 2: N = 0 (пограничный случай)
	t.Run("N=0", func(t *testing.T) {
		result := sliceDig(0)
		assert.Empty(t, result, "Слайс должен быть пустым")
	})

	// Тест 3: Проверка уникальности элементов
	t.Run("UniqueElements", func(t *testing.T) {
		n := 100
		result := sliceDig(n)
		unique := make(map[int]bool)
		for _, num := range result {
			if unique[num] {
				t.Fatalf("Найден дубликат: %d", num)
			}
			unique[num] = true
		}
	})
}

// Пример использования (для документации)
func ExampleSliceDig() {
	fmt.Println(sliceDig(3))
	// Output: [0 1 2]
}
