package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	summ := 0
	for i := 0; i < k; i++ {
		summ += nums[i]
	}
	curr := summ
	for i := k; i < len(nums); i++ {
		fmt.Println(i, nums[i], nums[i-k])
		curr += nums[i]
		curr -= nums[i-k]
		if curr > summ {
			summ = curr
		}
	}
	return float64(summ) / float64(k)
}
func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
