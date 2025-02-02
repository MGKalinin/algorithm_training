package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 36. Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/description/?envType=problem-list-v2&envId=29p0sxl6

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// line = strings.TrimSpace(line) // Убираем символ новой строки и пробелы

	// решить через многопоточность
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// Output: true

	fmt.Println(isValidSudoku(board))
	// Получаем количество горутин
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}

func isValidSudoku(board [][]byte) bool {
	var wg sync.WaitGroup
	ch := make(chan bool, 3)
	for i := 0; i < len(board); i++ {
		wg.Add(3)
		//проверка строки
		go func(row int) {
			defer wg.Done()
			if !checkUnique(board[row]) {
				ch <- false
			}
			ch <- true
		}(i)
		// проверка столбца
		go func(col int) {
			defer wg.Done()
			column := make([]byte, 9)
			for j := 0; j < 9; j++ {
				column[i] = board[col][j]
				if !checkUnique(column) {
					ch <- false
				}
				ch <- true
			}
		}(i)
		//	проверка блоков 3*3

	}

	return true
}

// написать функцию проверки byte слайса
func checkUnique(items []byte) bool {
	checkmap := make(map[byte]bool)
	for _, v := range items {
		if v != '.' {
			if checkmap[v] {
				return false
			}
			checkmap[v] = true
		}
	}
	return true
}
