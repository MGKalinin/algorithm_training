package main

import (
	"fmt"
	"log"
	"math"
)

// B. Кольцевая линия метро https://contest.yandex.ru/contest/28730/problems/B/

func main() {
	// * число --
	var N, i, j int64
	_, err := fmt.Scan(&N, &i, &j)
	if err != nil {
		log.Fatal(err)
	}
	first := math.Abs(float64(i - j))
	one := math.Abs(float64(N - i))
	second := math.Abs(float64(N - j))
	three := one + second
	fmt.Println(first)
	fmt.Println(three)

}
