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
	reader := bufio.NewReader(os.Stdin)
	// считать размер и количество запросов
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line) //получается слайс строк без пробелов и символов новой строки
	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])
	// чтение массива и построение префиксных сумм
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		num, _ := strconv.Atoi(parts[i-1])
		prefixSum[i] = prefixSum[i-1] + num
	}
	//подсчёт префиксных сумм в заданном диапазоне
	var result strings.Builder
	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		sum := prefixSum[r] - prefixSum[l-1]
		result.WriteString(strconv.Itoa(sum) + "\n")

	}
	fmt.Print(result.String())
}
