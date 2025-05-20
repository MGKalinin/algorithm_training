package main

import "fmt"

func maxOperations(nums []int, k int) int {
	count := 0
	var temp []int
	copy(temp, nums) //копия для удаления элементов || указатель?
	start, end := 0, len(temp)-1
	case:
		start<end{
			if k-temp[start]!=temp[end]{

	}
	}
	return count
}
func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}
