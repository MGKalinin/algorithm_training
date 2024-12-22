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
	for i := 0; i <= N; i++ {
		totalStations = append(totalStations, i)
	}
	betweenStations := totalStations[i:j]
	betweenStationsCircle1 := totalStations[:i]
	betweenStationsCircle2 := totalStations[j:]

	for _, num := range betweenStationsCircle1 {
		betweenStationsCircle2 = append(betweenStationsCircle2, num)
	}
	fmt.Println(totalStations)
	fmt.Println(betweenStations)
	fmt.Println(betweenStationsCircle1)

}
