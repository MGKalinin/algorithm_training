package main

import "fmt"

// Largest Prime Factor https://projecteuler.net/problem=3

func main() {
	/* Простые множители числа 13195 — это 5 , 7 , 13 и 29
	Какой самый большой простой множитель числа 600851475143 ? */
	findResult(13195)
}

// функция нахождения простых чисел и простых множителей,
// и возвращает максимальный простой множитель числа
func findResult(n int) int {
	prNum := []int{}
	for i := 0; i*i < n; i++ {
		if i <= 1 {
			continue
		}
		if i <= 3 {
			prNum = append(prNum, i)
		}
		if i%2 == 0 || i%3 == 0 {
			continue
		} else {
			prNum = append(prNum, i)
		}
	}
	fmt.Println(prNum)
	prDelitel := []int{}
	for _, num := range prNum {
		for n % num {
			prDelitel = append(prDelitel, num)

		}
	}
}
