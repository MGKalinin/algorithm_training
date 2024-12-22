package main

import (
	"fmt"
	"log"
)

// B. Кольцевая линия метро https://contest.yandex.ru/contest/28730/problems/B/
// 10 1 9
func main() {
	// * число --
	var N, i, j int
	_, err := fmt.Scan(&N, &i, &j)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(N)
	// fmt.Println(i)
	// fmt.Println(j)
	a := 0
	for k := i + 1; k < j; k++ {
		// fmt.Println(a)
		a++ //между станциями
	}
	// fmt.Println(a)
	b := 0
	for k := 1; k < i; k++ {
		b++ //до i
	}
	// fmt.Println(b)

	for k := j; k < N; k++ {
		b++ //от j до конца
	}
	// fmt.Println(b)
	if a < b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
