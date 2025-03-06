package main

import (
	"fmt"
	"runtime"
	"strings"
)

// 674. Longest Continuous Increasing Subsequence
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/?envType=problem-list-v2&envId=27tz3m7t

func main() {
	nums := []int{1, 3, 5, 4, 7}
	//Output: 3

	fmt.Println(findLengthOfLCIS(nums))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i <= len(dp); i++ {

	}
}
