package main

import "fmt"

// 238. Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность

	nums := []int{1, 2, 3, 4}
	// Output: [24,12,8,6]

	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {

}
