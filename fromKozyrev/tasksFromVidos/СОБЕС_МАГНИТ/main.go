package main

import "fmt"

//1. Смерджить 2 упорядоченных массива в 1 упорядоченный массив.
//Пример:
//дано: [1, 2, 3, 4, 5], [4, 5, 6, 7,8]
//Результат: [1, 2, 3, 4, 4, 5, 5, 6, 7, 8]

//func merge(a []int, b []int) []int {
//}

// 2. Смерджить произвольное количество каналов в один канал
// TODO для задачи слияния произвольного количества каналов в один паттерн Fan-In идеально подходит
func merge(chList ...chan int) chan int {

}

func main() {
	a := make(chan int, 2)
	a <- 1
	a <- 2
	close(a)

	b := make(chan int, 1)
	b <- 1
	close(b)

	res := merge(a, b)
	for v := range res {
		fmt.Println(v)
	}
}
