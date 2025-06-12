package main

import "fmt"

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	maxCount := 0
	currCount := 0
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currCount++
		}
	}
	maxCount = currCount
	for i := 1; i <= len(s)-k; i++ {
		if vowels[s[i-1]] {
			currCount--
		}
		if vowels[s[i+k-1]] {
			currCount++
		}
		if currCount > maxCount {
			maxCount = currCount
		}
	}
	return maxCount
}
func main() {
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels(s, k))
}
