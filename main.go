package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	s := "acb"    //искомое
	t := "ahbgdc" //где ищем
	// Output: true
	// поделить искомую подпоследовательность на два-запустить две горутины

	fmt.Println(isSubsequence(s, t))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func isSubsequence(s string, t string) bool {
	start := time.Now()

	res := make([]string, len(s))
	type indexStore struct {
		ind int
		val string
	}
	ch := make(chan indexStore, len(s))

	var wg sync.WaitGroup
	wg.Add(100)

	go func() {
		defer wg.Done()
		for _, char := range t {
			for k, j := range s {
				if j == char {
					ch <- indexStore{ind: k, val: string(j)}
				}
			}
		}
		close(ch)
	}()

	go func() {
		wg.Wait()
	}()

	for item := range ch {
		res[item.ind] = item.val
	}

	// Сравнение полученного результата с исходной строкой s
	result := ""
	for _, val := range res {
		result += val
	}
	// fmt.Println("Result:", result)

	// Замер времени выполнения
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
	return result == s

}
