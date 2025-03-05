package main

import (
	"fmt"
	"runtime"
	"strings"
)

// 151. Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	s := "the sky is blue"
	//Output: "blue is sky the"

	fmt.Println(reverseWords(s))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func reverseWords(s string) string {
	fields := strings.Fields(s)         // Разделяет строку на слова
	result := strings.Join(fields, " ") // Объединяет слова с одним пробелом
	start, end := 0, len(s)-1
	for start < end {
		if result[start] != ' ' || result[end] != ' ' {

		}
	}
	return s
}
