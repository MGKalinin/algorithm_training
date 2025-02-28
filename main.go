package main

import (
	"fmt"
	"runtime"
)

// 395. Longest Substring with At Least K Repeating Characters
// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	s := "aaabb"
	k := 3
	//Output: 3
	//Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.
	fmt.Println(longestPalindrome(s))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func longestSubstring(s string, k int) int {
	//Вот пошаговый алгоритм:
	//Подсчитать частоту символов в строке.
	//	Найти символы, которые встречаются реже, чем k. Они разделят строку на подстроки.
	//	Запустить горутину для обработки каждой подстроки параллельно.
	//	Использовать каналы для сбора результатов.
	//	Вернуть максимальную длину подходящей подстроки.

}
