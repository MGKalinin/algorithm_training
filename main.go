package main

import "fmt"

// 605. Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 1
	// Output: true
	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 2
	// Output: false
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2
	// false
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			n -= 1
		} else {
			continue
		}
	}
	if n == 0 {
		return true
	} else {
		return false
	}
}
