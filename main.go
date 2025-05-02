package main

import (
	"fmt"
	"runtime"
	//"strings"
)

// 238. Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	nums := []int{1, 2, 3, 4}
	//Output: [24,12,8,6]

	fmt.Println(productExceptSelf(nums))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

// productExceptSelf Product of Array Except Self
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefix := make([]int, n)
	prefix[0] = 1
	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		//fmt.Println("prefix", prefix, prefix[i-1], nums[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
		//fmt.Println("suffix", suffix, suffix[i+1], nums[i+1])
	}
	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
		//fmt.Println(result)
	}
	return result
}
