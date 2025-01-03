package main

import (
	"bufio"
	"fmt"
	"os"
)

// D. Правильная, круглая, скобочная https://contest.yandex.ru/contest/29075/problems/D/

func main() {
	reader := bufio.NewReader(os.Stdin)
	// считать размер массива
	line, _ := reader.ReadString('\n')
	// parts := strings.Fields(line) //получается слайс строк без пробелов и символов новой строки
	// N, _ := strconv.Atoi(parts[0]) //N групп
	// M, _ := strconv.Atoi(parts[1]) //M аудиторий
	// // чтение массива
	// line, _ = reader.ReadString('\n')
	// parts = strings.Fields(line)
	// groupN := make([]int, N) //слайс челов по группам
	// for i := 0; i < N; i++ {
	// 	groupN[i], _ = strconv.Atoi(parts[i])
	// }
	fmt.Println(line)
	line = line[:len(line)-1]
	for _, ch := range line {
		if ch == '(' {
			fmt.Println(ch)
		}
	}
}
