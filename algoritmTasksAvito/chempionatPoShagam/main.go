package main

import (
	"fmt"
	"reflect"
)

//Мы в Авито любим проводить соревнования, — недавно мы устроили чемпионат
//по шагам. И вот настало время подводить итоги!
//Необходимо определить userIds участников, которые прошли наибольшее
//количество шагов steps за все дни, не пропустив ни одного дня соревнований.
//Пример
//# Пример 1
//# ввод
//statistics = [
//[{ userId: 1, steps: 1000 }, { userId: 2, steps: 1500 }],
//[{ userId: 2, steps: 1000 }]
//]
//# вывод
//champions = { userIds: [2], steps: 2500 }
//# Пример 2
//statistics = [
//[{ userId: 1, steps: 2000 }, { userId: 2, steps: 1500 }],
//[{ userId: 2, steps: 4000 }, { userId: 1, steps: 3500 }]
//]
//# вывод
//champions = { userIds: [1, 2], steps: 5500 }

type Statistic struct {
	UserID int
	Steps  int
}
type Result struct {
	UserIDs []int
	Steps   int
}

func getChampions(statistics [][]Statistic) Result {
	// Реализуй функцию
}
func main() {

}
