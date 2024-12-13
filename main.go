package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	// Читаем строки до конца ввода или до определенного количества строк
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	// Выводим прочитанные строки
	for _, line := range lines {
		fmt.Print(line)
	}
}
