package main

import (
	"fmt"
	"runtime"
)

// 1641. Count Sorted Vowel Strings
// https://leetcode.com/problems/count-sorted-vowel-strings/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	n := 2
	//Output: 15
	//Explanation: The 15 sorted strings that consist of vowels only are
	//["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
	//	Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.

	fmt.Println(countVowelStrings(n))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func countVowelStrings(n int) int {

}
