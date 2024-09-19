// A. Двоичный поиск
package main

import (
	"fmt"
	"log"
)

// Функция считывает данные
func main() {
	// Считываем первую строку
	var nk = make([]int, 2)
	for i := 0; i < len(nk); i++ {
		_, err := fmt.Scan(&nk[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(nk)
	N := nk[0] // Длина первого массива
	K := nk[1] // Длина второго массива
	// Считываем первый массив
	var arr1 = make([]int, N)
	for i := 0; i < len(arr1); i++ {
		_, err := fmt.Scan(&arr1[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(arr1)
	// Считываем второй массив
	var arr2 = make([]int, K)
	for i := 0; i < len(arr2); i++ {
		_, err := fmt.Scan(&arr2[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(N, K, arr1, arr2)
}
