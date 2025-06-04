package main

import "fmt"

func longestOnes(nums []int, k int) int {
	l, r, cnt, maxLen := 0, 0, 0, 0
	for r < len(nums) {
		if nums[r] == 0 {
			cnt++
		}
		for cnt > k {
			if nums[l] == 0 {
				cnt--
			}
			l++
		}
		maxLen = max(maxLen, r-l+1)
		r++
	}
	return maxLen
}
func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println(longestOnes(nums, k))
}
