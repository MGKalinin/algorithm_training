package main

import "fmt"

/*
	Дано 2 отсортированных по возрастанию массива одинаковой длины.

Требуется получить результирующий массив,
содержащий все уникальные элементы двух массивов в порядке возрастания.
medium
[1,3,5], [3,4,5] => [1,3,4,5]
используем 2 указателя
какая алгоритмическая сложность у задачи?
*/
func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	f, s := 0, 0
	for f < len(a) && s < len(b) {
		switch {
		case a[f] < b[s]:
			if len(result) == 0 || a[f] != result[len(result)-1] {
				result = append(result, a[f])
			}
			f++

		case a[f] > b[s]:
			if len(result) == 0 || b[s] != result[(len(result)-1)] {
				result = append(result, b[s])
			}
			s++

		default:
			if len(result) == 0 || a[f] != result[len(result)-1] {
				result = append(result, a[f])
			}
			f++
			s++
		}
	}

	//add last elements
	for f < len(a) {
		if len(result) == 0 || a[f] != result[len(result)-1] {
			result = append(result, a[f])
		}
		f++
	}

	for s < len(b) {
		if len(result) == 0 || b[s] != result[(len(result)-1)] {
			result = append(result, b[s])
		}
		s++
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	fmt.Println(merge(a, b))
}
