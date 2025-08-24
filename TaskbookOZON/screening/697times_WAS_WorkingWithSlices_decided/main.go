package main

import "fmt"

// Что выведет данная программа. Почему так ?
func main() {
	nums := []int{1, 2, 3} //len 3 cap 3
	addNum(nums[0:2])
	fmt.Println(nums) // ? 1 2 -> 1 2 4
	addNums(nums[0:2])
	fmt.Println(nums) // ? 1 2 -> будет новый слайс (так как nums cap 3) со ссылкой на новый базовый массив 1 2 5 6
	// но в main при вызове отобразится 1 2 4 и 1 2 4
}
func addNum(nums []int) {
	nums = append(nums, 4)
}
func addNums(nums []int) {
	nums = append(nums, 5, 6)
}
