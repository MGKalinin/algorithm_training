package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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
	year, _ := strconv.Atoi(parts[2])
	// форматирование даты в нужный формат
	formattedDateDMY := fmt.Sprintf("%02d-%02d-%04d", day, month, year)
	formattedDateMDY := fmt.Sprintf("%02d-%02d-%04d", month, day, year)

	// Layout-ы для парсинга
	layoutMDY := "01-02-2006" // Месяц-День-Год
	layoutDMY := "02-01-2006" // День-Месяц-Год

	// Проверка даты в формате Месяц-День-Год
	_, errMDY := time.Parse(layoutMDY, formattedDateMDY)
	// Проверка даты в формате День-Месяц-Год
	_, errDMY := time.Parse(layoutDMY, formattedDateDMY)

	if errMDY == nil || errDMY == nil {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
