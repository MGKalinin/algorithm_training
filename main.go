package main

import (
	"fmt"
	"runtime"
)

// 746. Min Cost Climbing Stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	cost := []int{10, 15, 20}
	//Output: 15
	fmt.Println(minCostClimbingStairs(cost))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}
func minCostClimbingStairs(cost []int) int {

}
