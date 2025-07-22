package main

import "fmt"

// Что выведет данная программа. Почему так ?
func main() {
	nums := []int{1, 2, 3}
	addNum(nums[0:2])
	fmt.Println(nums) // 1 2 4- slice  1 2- addNum добавляет 4- всё норм -capacity позволяет-вывод 1 2 4
	addNums(nums[0:2])
	fmt.Println(nums) // 1 2 4 - slice 1 2 4- addNums- берёт 1 2  и добавляет 5 6 -не хватает capacity-аллокация, но
	// но в main ссылка на прежний slice 1 2 4- вывод будет 1 2 4
}
func addNum(nums []int) {
	nums = append(nums, 4)
}
func addNums(nums []int) {
	nums = append(nums, 5, 6)
}
