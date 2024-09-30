package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 99. Двоичный поиск
	// * число --
	var N, K int // N элементов первого массива, K элементов второго массива
	_, err := fmt.Scan(&N, &K)
	if err != nil {
		log.Fatal(err)
	}

	// первый массив
	reader := bufio.NewReader(os.Stdin)
	// var n int
	// fmt.Fscan(reader, &N)
	first := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &first[i])
	}
	// второй массив
	second := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscan(reader, &second[i])
	}
	// fmt.Println(first, second)
	// если K встречается в N- "YES"
	// бинарный поиск
	for i := range second {
		if binSearch(first[:], second[i]) { // преобразовать массив в срез?
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func binSearch(arr []int, target int) bool {
	start := 0
	finish := len(arr) - 1
	for start <= finish {
		mid := (start + finish) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			start = mid + 1
		} else {
			finish = mid - 1
		}
	}
	return false
}
