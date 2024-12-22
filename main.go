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
	totalStations := []int{} //всего станций
	for i := 1; i <= N; i++ {
		totalStations = append(totalStations, i)
	}
	// fmt.Println(totalStations)
	between := totalStations[i : j-1]
	a := len(between)
	// fmt.Println(between)
	parts := append(totalStations[:i-1], totalStations[j:]...)

	// fmt.Println(parts)
	b := len(parts)
	if a > b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}
