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
	//fmt.Println(dp)
	dp[1] = prices[0]
	//fmt.Println(dp)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] - prices[i-1]
	}
	fmt.Println(dp)
	return dp[n]
}

func maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
