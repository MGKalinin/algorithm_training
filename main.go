package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// A. Префиксные суммы https://contest.yandex.ru/contest/29075/problems/A/

func main() {
	var size, quantity int // размер массива ,количество запросов
	_, err := fmt.Scan(&size, &quantity)
	if err != nil {
		log.Fatal(err)
	}
	//сам массив
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < quantity; i++ { //считать пары чисел в количесвте запросов
		scanner.Scan()
		line := scanner.Text()
		part := strings.Split(line, " ")
		l, _ := strconv.Atoi(part[0]) //старт подсчёта
		r, _ := strconv.Atoi(part[1]) //финиш подсчёта
		//вот здесь надо подсчитать сумму и вывести в печать
		sum := 0
		for j := l - 1; j < r; j++ {
			sum += arr[j]
		}
		fmt.Println(sum)
	}

}
