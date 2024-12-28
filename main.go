package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// C. Даты https://contest.yandex.ru/contest/28730/problems/C/

// 1 2 2003
// 3 3 2067 // тест 3 fail

func main() {
	// Чтение строки из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ans := scanner.Text()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(ans, " ")
	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	// year, _ := strconv.Atoi(parts[2])
	if day <= 12 && month <= 12 && day != month {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}

}
