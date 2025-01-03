package main

import (
	"bufio"
	"os"
	"strings"
)

// E. Сумма трёх https://contest.yandex.ru/contest/29075/problems/E/

func main() {
	reader := bufio.NewReader(os.Stdin)
	// считать размер массива
	line, _ := reader.ReadString('\n')

	line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

}
