package main

import (
	"fmt"
	"log"
)

// D. Строительство школы https://contest.yandex.ru/contest/28730/problems/D/

func main() {
	// * число --
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	// * массив чисел известной длины --
	var arr = make([]int, number)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	mid := len(arr) / 2
	fmt.Println(arr[mid])
}
