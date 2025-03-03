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

}
