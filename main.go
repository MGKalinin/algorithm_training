package main

import "fmt"

// 238. Product of Array Except Self
func main() {
	nums := []int{1, 2, 3, 4}
	n := len(nums)
	result := make([]int, n)
	prefix := prefixProduct(nums)
	suffix := suffixProduct(nums)
	for i := 0; i < n; i++ {
		if i == 0 {
			result[i] = suffix[i+1]
		} else if i == n-1 {
			result[n-1] = prefix[i-1]
		} else {
			result[i] = prefix[i-1] * suffix[i+1]
		}
	}
	fmt.Println(result)
}

// функция вычисляет перфиксное произведение
func prefixProduct(nums []int) []int {
	n := len(nums)
	prefixProd := make([]int, n)
	prefixProd[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixProd[i] = prefixProd[i-1] * nums[i]
	}
	// fmt.Println(prefix_prod)
	return prefixProd
}

// функция вычисляет суффиксное произведение
func suffixProduct(nums []int) []int {
	n := len(nums)
	suffixProd := make([]int, n)
	suffixProd[n-1] = nums[n-1]
	for i := (n - 2); i >= 0; i-- {
		suffixProd[i] = nums[i] * suffixProd[i+1]
	}
	return suffixProd
}
