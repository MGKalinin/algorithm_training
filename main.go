package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A. Префиксные суммы https://contest.yandex.ru/contest/29075/problems/A/

func main() {
	// * строку --n и q
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0]) //размер массива
	q, _ := strconv.Atoi(parts[1]) //количество запросов
	//массив размера n
	// * массив чисел известной длины --
	var arr = make([]int, n)
	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, " ")
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(parts[i])
	}
	// Вычисление префиксных сумм
	prefixSums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSums[i] = prefixSums[i-1] + arr[i-1]
	}
	// fmt.Println("\n", prefixSums)
	// fmt.Println(n, q)
	// fmt.Println(arr)
	//количество запросов l, r
	for i := 0; i < q; i++ {
		scanner.Scan()
		line = scanner.Text()
		parts = strings.Split(line, " ")
		l, _ := strconv.Atoi(parts[0]) //левая граница отрезка
		r, _ := strconv.Atoi(parts[1]) //правая граница отрезка
		// fmt.Println("---", l, r)
		// Подсчет суммы на отрезке [l, r]
		sum := prefixSums[r] - prefixSums[l-1]
		fmt.Println(sum)

	}
}

// функция подсчёта префиксной суммы в указанном диапазоне
// прикрутить тестирование функции prefixSum?+
// func prefixSum(arr []int, n, l, r int) int {
// 	sum := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		sum[i] = sum[i-1] + arr[i-1]
// 	}
// 	return sum[r] - sum[l-1]
// }
