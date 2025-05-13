package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//Output: 49
	result := maxArea(height)
	fmt.Println("Output:", result)
	//fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	n := len(height)
	maxVal := 0          //максимальный объём
	start, end := 0, n-1 //левый, правый край контейнера
	for start < end {
		temp := (minn(height[start], height[end])) * (end - start)
		//fmt.Println(temp)
		if temp > maxVal {
			maxVal = temp
		}
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	//fmt.Println(maxVal)
	return maxVal
}

// minn возвращает минимальное значение
func minn(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
