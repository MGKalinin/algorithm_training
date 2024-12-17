package main

import (
	"fmt"
	"log"
)

// B. Кольцевая линия метро https://contest.yandex.ru/contest/28730/problems/B/

func main() {
	// * число --
	var N, i, j int64
	_, err := fmt.Scan(&N, &i, &j)
	if err != nil {
		log.Fatal(err)
	}
	first := j - i
	one := (N - i)
	second := (N - j)
	fmt.Println(first)
	fmt.Println(one)
	fmt.Println(second)

}
