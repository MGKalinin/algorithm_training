package main

import (
	"fmt"
	"runtime"
	"strings"
)

// 416. Partition Equal Subset Sum
// https://leetcode.com/problems/partition-equal-subset-sum/description/?envType=problem-list-v2&envId=27tz3m7t

func main() {
	nums := []int{1, 5, 11, 5}
	//true

	fmt.Println(canPartition(nums))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func canPartition(nums []int) bool {

}
