package main

import (
	"fmt"
	"runtime"
)

// 5. Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	s := "babad"
	//Output: "bab"
	//Explanation: "aba" is also a valid answer.

	fmt.Println(longestPalindrome(s))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func longestPalindrome(s string) string {

}
