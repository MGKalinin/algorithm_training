package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Требуется реализовать функцию uniqRandn, которая генерирует слайс длины n уникальных,
//рандомных чисел.

func main() {
	fmt.Println(uniqRandn(10))
}

func uniqRandn(n int) []int {
	result := make([]int, 0, n)      //инициализация slice длиной 0 и ёмкость n; инициализация slice без 0
	rand.Seed(time.Now().UnixNano()) // инициализация генератора
	unic := make(map[int]struct{})
	for len(result) < n {
		r := rand.Int()            // случайное int
		if _, ok := unic[r]; !ok { // value, ok
			unic[r] = struct{}{}
			result = append(result, r)
		}
	}
	return result
}

// https://go.dev/play/p/ggJxkffaR8j
