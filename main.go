package main

import (
	"fmt"
	"math"
	"runtime"
)

// 334. Increasing Triplet Subsequence
// https://leetcode.com/problems/increasing-triplet-subsequence/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	nums := []int{2, 4, -2, -3}
	//Output: true
	//Explanation: Any triplet where i < j < k is valid.

	fmt.Println(increasingTriplet(nums))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

// increasingTriplet Increasing Triplet Subsequence
func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt
	for _, num := range nums {
		if num > second {
			return true
		}
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		}
	}
	return false
}
