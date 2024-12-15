package main

import "fmt"

// Largest Prime Factor https://projecteuler.net/problem=3

func main() {
	/* Простые множители числа 13195 — это 5 , 7 , 13 и 29
	Какой самый большой простой множитель числа 600851475143 ? */
	findRes(600851475143)
}

// функция нахождения простых чисел и простых множителей,
// и возвращает максимальный простой множитель числа
func primeNumber(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	return true
}

func findRes(n int) {
	ans := []int{}
	for i := 2; i <= n; i++ {
		if primeNumber(i) && n%i == 0 {
			ans = append(ans, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	fmt.Println(ans)
	maxelem := ans[0]
	for _, num := range ans {
		if num > maxelem {
			maxelem = num
		}
	}
	fmt.Println(maxelem)
}
