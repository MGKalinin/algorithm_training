package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// A. Количество равных максимальному https://contest.yandex.ru/contest/28738/problems/A/

func main() {
	arr := []int{}
	// * строку --
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Ошибка преобразования '%s' в число: %v, пропускаем.", line, err)
			continue
		}
		if num == 0 {
			break
		}
		arr = append(arr, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения из stdin: %v", err)
	}
	fmt.Println(arr)
	// посчитать к-во чисел равных максимальному
}
