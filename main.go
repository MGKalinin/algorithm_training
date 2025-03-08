package main

import (
	"fmt"
	"runtime"
	"strings"
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

func productExceptSelf(nums []int) []int {

}
