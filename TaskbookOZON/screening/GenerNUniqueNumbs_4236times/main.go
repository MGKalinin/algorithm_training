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
	result := make([]int, 0, n)                            //инициализация slice длиной 0 и ёмкость n; инициализация slice без 0
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // инициализация генератора
	unic := make(map[int]struct{})
	for len(result) < n {
		r := rng.Int()             // случайное int
		if _, ok := unic[r]; !ok { // value, ok
			unic[r] = struct{}{}
			result = append(result, r)
		}
	}
	return result
}
