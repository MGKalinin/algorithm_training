package main

import (
	"fmt"
	"runtime"
)

// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	n := 4
	//Output: 3
	fmt.Println(fib(n))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func fib(n int) int {

}
