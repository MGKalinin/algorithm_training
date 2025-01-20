package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	s := "abc"    //искомое
	t := "ahbgdc" //где ищем
	// Output: true
	// поделить искомую подпоследовательность на два-запустить две горутины

	fmt.Println(isSubsequence(s, t))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func isSubsequence(s string, t string) bool {
	mid := len(t) / 2
	// fmt.Println(t[:mid])
	// fmt.Println(t[mid:])
	res := make([]string, len(t[:mid]))
	type indexStore struct {
		ind int
		val string
	}
	ch1 := make(chan indexStore, len(t[:mid])) //для 1-й половины
	// ch2 := make(chan rune, len(t[mid:])) //для 2-й половины

	var wg sync.WaitGroup

	for _, char := range t[:mid] { //идём по первой половине
		wg.Add(1)
		go func(char rune) {
			defer wg.Done()
			for k, j := range s { //искомое
				if j == char {
					ch1 <- indexStore{ind: k, val: string(j)}
				}

			}
		}(char)
	}
	wg.Wait()
	go func() {
		close(ch1)
	}()

	for item := range ch1 {
		res[item.ind] = item.val
	}
	fmt.Println(res)
	return true
}
