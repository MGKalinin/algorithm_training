package main

import (
	"fmt"
	"runtime"
)

// 5. Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	s := "babad"
	//Output: "bab"
	fmt.Println(longestPalindrome(s))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func longestPalindrome(s string) string {

}
