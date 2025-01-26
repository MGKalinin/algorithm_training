package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// 1704. Determine if String Halves Are Alike
// https://leetcode.com/problems/determine-if-string-halves-are-alike/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	s := "textbook"
	// Output: true

	fmt.Println(halvesAreAlike(s))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func halvesAreAlike(s string) bool {
	// start := time.Now()
	vowels := "aeiouAEIOU"
	count1 := 0
	count2 := 0
	mid := len(s) / 2
	var mux sync.Mutex
	var wg sync.WaitGroup
	for _, i := range s[:mid] {
		wg.Add(1)
		go func(i rune) {
			defer wg.Done()
			if strings.ContainsRune(vowels, i) {
				mux.Lock()
				count1++
				mux.Unlock()
			}
		}(i)
	}

	for _, i := range s[mid:] {
		wg.Add(1)
		go func(i rune) {
			defer wg.Done()
			if strings.ContainsRune(vowels, i) {
				mux.Lock()
				count2++
				mux.Unlock()
			}
		}(i)
	}
	wg.Wait()
	// elapsed := time.Since(start)
	// fmt.Printf("Execution time: %s\n", elapsed)

	return count1 == count2

}

func maximumWealth(accounts [][]int) int {
	// start := time.Now()

	var wg sync.WaitGroup
	n := len(accounts)
	res := make([]int, n)
	ch := make(chan int, n)
	for _, i := range accounts {
		// fmt.Println(i)
		wg.Add(1)
		go func(i []int) {
			wg.Done()
			summ := 0
			for _, k := range i {
				summ += k
			}
			ch <- summ
		}(i)
	}
	wg.Wait()
	go func() {
		close(ch)
	}()
	for i := range ch {
		res = append(res, i)
	}
	maxx := 0
	for _, i := range res {
		if i > maxx {
			maxx = i
		}
	}
	// Замер времени выполнения
	// elapsed := time.Since(start)
	// fmt.Printf("Execution time: %s\n", elapsed)

	// fmt.Println(maxx)
	return maxx
}
