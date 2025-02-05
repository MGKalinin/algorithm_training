package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 36. Valid Sudoku // TODO: solved
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
				return
			}
			ch <- true
		}(i)
		// проверка столбца
		go func(col int) {
			defer wg.Done()
			column := make([]byte, 9)
			for j := 0; j < 9; j++ {
				column[j] = board[j][col]
			}
			if !checkUnique(column) {
				ch <- false
				return
			}
			ch <- true
		}(i)
		//	проверка блоков 3*3
		go func(block int) {
			defer wg.Done()
			rowBase := (block / 3) * 3
			colBase := (block % 3) * 3
			subBox := make([]byte, 9)
			index := 0
			for r := rowBase; r < rowBase+3; r++ {
				for c := colBase; c < colBase+3; c++ {
					subBox[index] = board[r][c]
					index++
				}
			}
			if !checkUnique(subBox) {
				ch <- false
				return
			}
			ch <- true
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for result := range ch {
		if !result {
			return false
		}
	}
	return true
}

// функция проверки byte слайса
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
