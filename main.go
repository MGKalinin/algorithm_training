package main

import (
	"fmt"
	"runtime"
)

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=problem-list-v2&envId=dynamic-programming

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		if prices[i] < prices[i+1] {
			//dp[i]=
		}
	}
	return dp[n]
}

func maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
