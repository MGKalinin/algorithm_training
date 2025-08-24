package main

import "fmt"

// Является ли слайс монотонным?
// Монотонная функция - функция одной переменной, определённая на некотором
// подмножестве действительных чисел, которая либо везде (на области своего
// определения) не убывает, либо везде не возрастает.
// {1,7} - true
// {1,1} - true
// {3,3,1} - true
// {9,5,1} - true
// {23,5,23} - false
func monotonic(in []int) bool {
	isIncreasing, isDecreasing := true, true
	for i := 0; i < len(in)-1; i++ {
		if in[i] < in[i+1] {
			isDecreasing = false
		}
		if in[i] > in[i+1] {
			isIncreasing = false
		}
		if !isIncreasing && !isDecreasing {
			return false
		}
	}
	return isIncreasing || isDecreasing

}

func main() {
	a := []int{1, 7}      //- true
	b := []int{1, 1}      //- true
	c := []int{3, 3, 1}   //- true
	d := []int{9, 5, 1}   //- true
	e := []int{23, 5, 23} //- false
	fmt.Println(monotonic(a))
	fmt.Println(monotonic(b))
	fmt.Println(monotonic(c))
	fmt.Println(monotonic(d))
	fmt.Println(monotonic(e))
}
