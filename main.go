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
	totalStations := []int{} //всего станций
	for i := 1; i <= N; i++ {
		totalStations = append(totalStations, i)
	}
	between := totalStations[i:j]
	part1 := totalStations[:i]
	part2 := totalStations[j:]

	fmt.Println(totalStations)
	fmt.Println(between)
	fmt.Println(part1)
	fmt.Println(part2)

}
