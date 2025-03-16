package main

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	//"strings"
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
	sort.Ints(arr)
	pairs := [][]int{}
	minDiff := math.MaxInt64
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
			pairs = [][]int{{arr[i], arr[i+1]}}
		} else if diff == minDiff {
			pairs = append(pairs, []int{arr[i], arr[i+1]})
		}
	}
	return pairs
}

//TODO: переделать в многопоточность
