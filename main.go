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

	s := "abc"    //искомое
	t := "ahbgdc" //где ищем
	// Output: true
	// поделить искомую подпоследовательность на два-запустить две горутины

	fmt.Println(isSubsequence(s, t))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func isSubsequence(s string, t string) bool {
	start := time.Now()

	mid := len(t) / 2
	// fmt.Println(t[:mid])
	// fmt.Println(t[mid:])
	res := make([]string, len(s))
	type indexStore struct {
		ind int
		val string
	}
	ch1 := make(chan indexStore, len(t[:mid])) //для 1-й половины
	ch2 := make(chan indexStore, len(t[mid:])) //для 2-й половины

	var wg sync.WaitGroup
	wg.Add(2)
	// Горутина для первой половины
	go func() {
		defer wg.Done()
		for _, char := range t[:mid] {
			for k, j := range s {
				if j == char {
					ch1 <- indexStore{ind: k, val: string(j)}
				}
			}
		}
		close(ch1)
	}()

	// Горутина для второй половины
	go func() {
		defer wg.Done()
		for _, char := range t[mid:] {
			for k, j := range s {
				if j == char {
					ch2 <- indexStore{ind: k, val: string(j)}
				}
			}
		}
		close(ch2)
	}()
	wg.Wait()

	for item := range ch1 {
		res[item.ind] = item.val
	}
	for item := range ch2 {
		res[item.ind] = item.val
		fmt.Println(res)
	}
	return true
}

// "mdadm": "4.1-11"
//   "parted": "3.4-1"
// POST partnumbers-2025-01-21/_doc
// {
//     "@timestamp": "2025-01-21T12:00:03",
//     "AQPartNumber": "70079872",
//     "device_id": "test624",
//     "components": [
//       {
//         "AQPartNumber": "10088693",
//         "device_id": "ANS685V4000124"
//       }
//     ]
// }
