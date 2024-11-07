package main

import "fmt"

// 238. Product of Array Except Self
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix_sum := make([]int, n+1)
	fmt.Println(prefix_sum)
	prefix_sum[0] = 0
	for i := 1; i <= n; i++ {
		prefix_sum[i] = prefix_sum[i-1] + nums[i-1]
		fmt.Println(prefix_sum)
	}
	return prefix_sum
}

// функция вычисляет перфиксное произведение

// функция вычисляет суффиксное произведение
