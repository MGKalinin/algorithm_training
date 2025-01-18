package main

import (
	"fmt"
	"runtime"
)

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	s := "abc"
	t := "ahbgdc"
	// Output: true
	// поделить искомую подпоследовательность на два-запустить две горутины

	fmt.Println(isSubsequence(s, t))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func isSubsequence(s string, t string) bool {

}
