package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 70. Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	n := 3
	//Output: 3
	//Explanation: There are three ways to climb to the top.
	//	1. 1 step + 1 step + 1 step
	//	2. 1 step + 2 steps
	//	3. 2 steps + 1 step

	fmt.Println(climbStairs(n))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func climbStairs(n int) int {

}
