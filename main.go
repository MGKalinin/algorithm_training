package main

import "fmt"

// 1768. Merge Strings Alternately
// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')

	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы
	// fmt.Println(line)
	word1 := "abc"
	word2 := "pqr"
	// Output: "apbqcr"
	fmt.Println(mergeAlternately(word1, word2))
}
func mergeAlternately(word1 string, word2 string) string {
	var word3 string
	if len(word1) == len(word2) {
		for i, _ := range word1 {
			for k, _ := range word2 {
				// word3 = word1[i] + word2[k]
			}
		}
	}
	return word3
}
