package main

import (
	"fmt"
	"runtime"
)

// 1200. Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	arr := []int{4, 2, 1, 3}
	//Output: [[1,2],[2,3],[3,4]]
	//Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

	fmt.Println(minimumAbsDifference(arr))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func minimumAbsDifference(arr []int) [][]int {

}
