package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// B. Максимальная сумма https://contest.yandex.ru/contest/29075/problems/B/

func main() {
	reader := bufio.NewReader(os.Stdin)
	// считать размер массива
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line) //получается слайс строк без пробелов и символов новой строки
	n, _ := strconv.Atoi(parts[0])
	// чтение массива
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(parts[i])
	}
	// подсчёт префиксных сумм всего массива
	prefixSumm := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSumm[i] = prefixSumm[i-1] + arr[i-1]
	}
	// сравнить разницу минимальной префиксной суммы и максимальной префиксной суммы
maxSum:=
}

// 5
// 450402558 -840167367 -231820501 586187125 -627664644
