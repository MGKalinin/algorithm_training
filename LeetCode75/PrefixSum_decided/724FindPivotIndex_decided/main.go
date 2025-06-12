package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		right = total - nums[i] - left

		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
