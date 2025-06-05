package main

import "fmt"

func largestAltitude(gain []int) int {
	answer := make([]int, len(gain)+1)
	answer[0] = 0
	maxx := 0
	for i := 0; i < len(gain); i++ {
		answer[i+1] = answer[i] + gain[i]
		fmt.Println(answer)
		if answer[i+1] > maxx {
			maxx = answer[i+1]
		}
	}
	fmt.Println(answer)
	return maxx
}
func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}
