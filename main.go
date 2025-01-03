package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// D. Правильная, круглая, скобочная https://contest.yandex.ru/contest/29075/problems/D/

func main() {
	reader := bufio.NewReader(os.Stdin)
	// считать размер массива
	line, _ := reader.ReadString('\n')

	line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы
	stack := []rune{}
	for _, ch := range line {
		if ch == '(' {
			stack = append(stack, ch)

		} else if ch == ')' {
			if len(stack) == 0 {
				fmt.Println("NO")
				return
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		fmt.Println("YES")
		// return "YES"
	} else {
		fmt.Println("NO")
		// return "NO"
	}
}
