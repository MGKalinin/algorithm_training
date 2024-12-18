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
	fmt.Println(N)
	fmt.Println(i)
	fmt.Println(j)
	first := []int{} // общий список чисел
	// second := []int{} //списко от начала до i
	// three := []int{}  //список от j до конца
	for k := 1; k <= N; k++ {
		first = append(first, k)
	}
	fmt.Println(first)
	onePart := first[:i]

	fmt.Println(onePart)
	twoPart := first[j+1:]
	fmt.Println(twoPart)

}
