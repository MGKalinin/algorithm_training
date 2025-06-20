package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 1.написать код для выдачи массива уникальных случайных чисел
func numberGenerator(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("errors: n must >0")
	}
	if n == 0 {
		return []int{}, nil
	}
	unicNums := make(map[int]struct{})
	slice := make([]int, 0, n)
	for len(slice) < n {
		num := rand.Int() // Случайное целое (int)
		if _, ok := unicNums[num]; !ok {
			unicNums[num] = struct{}{}
			slice = append(slice, num)
		}
	}
	return slice, nil
}

//2.нужно написать http сервер, который обрабатывает get запросы и возвращает список файлов в указанной директории

//3.реализовать worker pool. есть 10 задач(функций) каждая засыпает на 1 секунду и
//выводит номер воркера который эту задачу исполнил, количество воркеров задаётся при запуске

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := 10
	fmt.Println(numberGenerator(n))
}
