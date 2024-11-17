package main

import (
	"fmt"
	"log"
)

// 92. Туризм
func main() {

	var N int //количество точек в цепи
	_, err := fmt.Scan(&N)
	if err != nil {
		log.Fatal(err)
	}
	var mountainRange [][]int

	// * cлайс слайсовы чисел  длины N координат горной цепи
	for i := 0; i < N; i++ {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			log.Fatal(err)
		}
		mountainRange = append(mountainRange, []int{x, y})
	}

	var M int // количество трасс
	_, err = fmt.Scan(&M)
	if err != nil {
		log.Fatal(err)
	}
	var numberTrails [][]int
	// слайс слайсов  номера точек начала и конца трассы
	for i := 0; i < M; i++ {
		var s, f int
		_, err := fmt.Scan(&s, &f)
		if err != nil {
			log.Fatal(err)
		}
		numberTrails = append(numberTrails, []int{s, f})
	}
	fmt.Println(N)
	fmt.Println(mountainRange)
	fmt.Println(M)
	fmt.Println(numberTrails)
}

// написать тесты?
