package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	numRows := 5
	//Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Println(generate(numRows))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}
func generate(numRows int) [][]int {

}
