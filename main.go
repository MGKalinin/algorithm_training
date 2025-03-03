package main

import (
	"fmt"
	"runtime"
)

// 345. Reverse Vowels of a String
// https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	s := "IceCreAm"
	//Output: "AceCreIm"
	//Explanation:
	//	The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

	fmt.Println(reverseVowels(s))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func reverseVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	start, end := 0, len(s)
if s[start] && s[end]
}
