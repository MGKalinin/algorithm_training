package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for _, num := range nums1 {
		map1[num] = true
	}
	for _, num := range nums2 {
		map2[num] = true
	}

	firstChannel := make(chan int, len(nums1))
	secondChannel := make(chan int, len(nums2))

	go func() {
		temp := make(map[int]bool)
		for num := range map1 {
			if !map2[num] && !temp[num] {
				firstChannel <- num
				temp[num] = true
			}
		}
		close(firstChannel)
	}()

	go func() {
		temp := make(map[int]bool)
		for num := range map2 {
			if !map1[num] && !temp[num] {
				secondChannel <- num
				temp[num] = true
			}
		}
		close(secondChannel)
	}()

	diff1 := []int{}
	for num := range firstChannel {
		diff1 = append(diff1, num)
	}

	diff2 := []int{}
	for num := range secondChannel {
		diff2 = append(diff2, num)
	}

	return [][]int{diff1, diff2}
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}
