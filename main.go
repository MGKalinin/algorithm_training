package main

import (
	"fmt"
)

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	s := "bb"     //искомое
	t := "ahbgdc" //где ищем
	// Output: false

	fmt.Println(isSubsequence(s, t))
	// Получаем количество горутин
	// fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func isSubsequence(s string, t string) bool {
	// start := time.Now()
	//одна горутина ищет первую половину s, вторая -вторую?
	ans := ""
	for _, v := range t {
		for _, z := range s {
			if z == v {
				ans += string(v)
			}
		}
	}

	fmt.Println(ans)
	// Замер времени выполнения
	// elapsed := time.Since(start)
	// fmt.Printf("Execution time: %s\n", elapsed)
	return ans == s

}

// 2025-01-22T10:00:00
// func newSeach(s string, t string) bool {

// }
