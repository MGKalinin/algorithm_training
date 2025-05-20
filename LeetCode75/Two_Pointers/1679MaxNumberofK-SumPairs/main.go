package main

import "fmt"

func maxOperations(nums []int, k int) int {
	count := 0
	var temp []int
	copy(temp, nums) //копия для удаления элементов || указатель?
	start, end := temp[0], temp[1]
	for start < end {
		if k-temp[end] == temp[start] {
			temp = append(temp[:end], temp[end+1:]...)
			temp = append(temp[:start], temp[start+1:]...)
			count++
		}
	}
	return count
}
func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}
