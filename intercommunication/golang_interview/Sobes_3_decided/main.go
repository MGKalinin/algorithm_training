package main

import (
	"fmt"
)

// ------------------------------------------------------------------------------------------------------------
// Задача 1
// 1. Написать функцию, которая принимает число N и возвращает слайс размера N с уникальными числами.
// 2. Идеи как тестировать функцию?
// ------------------------------------------------------------------------------------------------------------
func sliceDig(n int) []int {
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func main() {
	fmt.Println(sliceDig(10))
}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func sliceDig(n int) []int {
//	rand.Seed(time.Now().UnixNano())
//	unique := make(map[int]bool)
//	nums := make([]int, 0, n)
//
//	for len(nums) < n {
//		num := rand.Intn(n * 10) // Генерируем числа в диапазоне [0, n*10)
//		if !unique[num] {
//			unique[num] = true
//			nums = append(nums, num)
//		}
//	}
//
//	return nums
//}
//
//func main() {
//	fmt.Println(sliceDig(10)) // Пример вывода: [45 12 3 89 7 23 56 1 34 67]
//}
