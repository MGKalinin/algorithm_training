package main

func main() {
	nums := []int{0, 1, 0, 3, 12}
	//Output: [1,3,12,0,0]
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
	//fmt.Println(nums)
}
