package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	count := 0
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for start < end {
		summ := nums[start] + nums[end]
		if summ == k {
			count++
			start++
			end--
		} else if summ < k {
			start++
		} else {
			end--
		}
	}
	return count
}
func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}
