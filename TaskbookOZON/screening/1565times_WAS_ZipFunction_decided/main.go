package main

import "fmt"

// Реализовать версию функции zip которая сможет соединять произвольное количество
// слайсов
// Выводит слайс слайсов  с элементами поиндексно [элементы с индексом 0][...с индесом 1][...с индексом 2]
func main() {
	s1, s2, s3 := []int{1, 2, 3}, []int{4, 5, 6, 7, 8}, []int{9, 10, 11}
	fmt.Println(zip(s1, s2, s3)) // [1 4 9] [2 5 10] [3 6 11]
}

func zip(s ...[]int) [][]int {
	// Проверка на пустоту
	if len(s) == 0 {
		return [][]int{}
	}

	// Находим минимальную длину среди всех слайсов
	minSize := len(s[0])
	for _, slice := range s {
		if len(slice) < minSize {
			minSize = len(slice)
		}
	}

	result := make([][]int, minSize)
	// Добавляем элементы по индексам
	for i := 0; i < minSize; i++ {
		// Инициализируем слайс для каждого индекса
		result[i] = make([]int, len(s))
		for j, slice := range s {
			result[i][j] = slice[i]
		}
	}
	return result
}
