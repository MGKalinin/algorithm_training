package main

import "fmt"

// 238. Product of Array Except Self
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productPrefix(nums))
	fmt.Println(productSuffix(nums))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix_sum := make([]int, n+1)
	prefix_sum[0] = 0
	for i := 1; i <= n; i++ {
		prefix_sum[i] = prefix_sum[i-1] + nums[i-1]

	}
	fmt.Println(prefix_sum)
	return prefix_sum
}

// функция вычисляет перфиксное произведение
func productPrefix(nums []int) []int {
	n := len(nums)
	prefix_prod := make([]int, n+1)
	prefix_prod[0] = 1
	for i := 1; i <= n; i++ {
		prefix_prod[i] = prefix_prod[i-1] * nums[i-1]
	}
	fmt.Println(prefix_prod)
	return prefix_prod
}

// функция вычисляет суффиксное произведение
func productSuffix(nums []int) []int {
	n := len(nums)
	prefix_prod := make([]int, n+1)
	prefix_prod[0] = 1
	for i := 0; i <= n-1; i++ {
		prefix_prod[i] = prefix_prod[i+1] * nums[i+1]
	}
	fmt.Println(prefix_prod)
	return prefix_prod
}
