package main

import (
	"fmt"
)

// 1071. Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	str1 := "ABCABC"
	str2 := "ABC"
	// Output: "ABC"
	// str1 := "ABABAB"
	// str2 := "ABAB"
	// Output: "AB"
	fmt.Println(gcdOfStrings(str1, str2))
}
func gcdOfStrings(str1 string, str2 string) string {

}
