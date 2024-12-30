package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// B. Максимальная сумма https://contest.yandex.ru/contest/29075/problems/B/

func main() {
	reader := bufio.NewReader(os.Stdin)
	// считать размер и количество запросов
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line) //получается слайс строк без пробелов и символов новой строки
	n, _ := strconv.Atoi(parts[0])
	// чтение массива
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	// подсчёт префиксной суммы
	maxResult, _ := strconv.Atoi(parts[0])
	prefixSum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		num, _ := strconv.Atoi(parts[i-1])
		prefixSum[i] = prefixSum[i-1] + num
		if prefixSum[i] > maxResult {
			maxResult = prefixSum[i]

		} else {
			continue
		}
	}
	fmt.Println(maxResult)
}

// 5
// 450402558 -840167367 -231820501 586187125 -627664644
