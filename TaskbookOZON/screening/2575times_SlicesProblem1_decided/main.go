package main

import "fmt"

// Что выведет программа?
func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)  // 0 1 2 длина (len) = 3, ёмкость (cap) = 4 (ёмкость увеличивается по алгоритму Go: 0 → 1 → 2 → 4)
	y := append(x, 3) // 0 1 2 3
	z := append(x, 4) // 0 1 2 4 ;z видит [0, 1, 2, 4]
	fmt.Println(y, z) // 0 1 2 4  0 1 2 4 ;y ссылается на тот же базовый массив, что x и z
}
func main() {
	a()
}
