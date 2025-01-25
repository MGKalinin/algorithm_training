package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 1672. Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}

	fmt.Println(maximumWealth(accounts))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
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
