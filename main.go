package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// C. Даты https://contest.yandex.ru/contest/28730/problems/C/

func main() {
	// * строку --
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ans := scanner.Text()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ans)
	// Примеры строк дат
	dateMDY := "01 25 2010"
	dateDMY := "27 05 2011"
	// Проверка даты в формате Месяц-День-Год
	_, errMDY := time.Parse(dateMDY, ans)
	// Проверка даты в формате День-Месяц-Год
	_, errDMY := time.Parse(dateDMY, ans)
	if errMDY == nil {

	} else if errDMY == nil {

	} else {

	}
}
