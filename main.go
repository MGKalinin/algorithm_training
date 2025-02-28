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
	minVal, maxRes := prices[0], 0
	for i := 1; i < n; i++ {
		if maxRes < prices[i]-minVal {
			maxRes = prices[i] - minVal
		} else if prices[i] < minVal {
			minVal = prices[i]
		}
	}
	return maxRes
}
