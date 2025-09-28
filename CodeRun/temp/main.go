package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 60. Кубики
func main() {

	// открыть файл
	file, err := os.Open("/Users/maksimkalinin/GolandProjects/algorithm_training/CodeRun/temp/input.txt")
	if err != nil {
		fmt.Println("can`t open file")
		return
	}
	defer file.Close()

	// создать ридер
	reader := bufio.NewReader(file)
	// создать врайтер
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	// продуть буффер
	defer writer.Flush()

	// прочитать первую строку
	line, _ := reader.ReadString('\n')
	ab := strings.Fields(strings.TrimSpace(line))
	// присвоить первое значение a, второе b
	a, _ := strconv.Atoi(ab[0])
	b, _ := strconv.Atoi(ab[1])

	// считать цвета а в слайс
	aColor := make([]int, 0, a) //len 0 чтоб нули не написал по умолчанию, сap a чтоб без аллокации памяти
	for dig := 0; dig < a; dig++ {
		line, _ := reader.ReadString('\n')
		val, _ := strconv.Atoi(strings.TrimSpace(line))
		aColor = append(aColor, val)
	}
	// считать цвета b в слайс
	bColor := make([]int, 0, b)
	for dig := 0; dig < b; dig++ {
		line, _ := reader.ReadString('\n')
		val, _ := strconv.Atoi(strings.TrimSpace(line))
		bColor = append(bColor, val)
	}
	// пройти двумя указателями по слайсу a & b -найти совпадения?
	// слайс общих цветов
	common := make([]int, 0) // здесь можно заебаться и написать функцию получения большего значения cap...
	// сортируем слайсы
	sort.Ints(aColor)
	sort.Ints(bColor)
	// два указателя
	first, second := 0, 0
	for first < len(aColor) && second < len(bColor) {
		if aColor[first] == bColor[second] {
			common = append(common, aColor[first])
			first++
			second++
		} else if aColor[first] < bColor[second] {
			first++
		} else {
			second++
		}
	}
	//fmt.Println(common)
	// в печать общие цвета
	if len(common) > 0 {
		writer.WriteString(strconv.Itoa(len(common)) + "\n")
		for i, val := range common {
			if i > 0 {
				writer.WriteString(" ") // чтобы не ставить пробел перед первым элементом и
				// не ставить пробел после последнего элемента
			}
			writer.WriteString(strconv.Itoa(val))
		}
		writer.WriteString("\n")
	} else {
		writer.WriteString("0\n") // ← здесь печатаем "0" и переводим строку
	}

	// уникальные цвета a
	f, s := 0, 0
	onlyA := make([]int, 0, len(aColor))
	for f < len(aColor) && s < len(bColor) {
		if aColor[f] == bColor[s] {
			f++
			s++
		} else if aColor[f] < bColor[s] {
			onlyA = append(onlyA, aColor[f])
			f++
		} else {
			s++
		}
	}
	// добавляем оставшиеся
	for f < len(aColor) {
		onlyA = append(onlyA, aColor[f])
		f++
	}

	// уникальные цвета b
	one, two := 0, 0
	onlyB := make([]int, 0, len(bColor))
	for one < len(aColor) && two < len(bColor) {
		if aColor[one] == bColor[two] {
			one++
			two++
		} else if aColor[one] < bColor[two] {
			one++
		} else {
			onlyB = append(onlyB, bColor[two])
			two++
		}
	}
	// добавляем оставшиеся
	for two < len(bColor) {
		onlyB = append(onlyB, bColor[two])
		two++
	}
	// в печать уникальные a
	if len(onlyA) > 0 {
		writer.WriteString(strconv.Itoa(len(onlyA)) + "\n")
		for i, val := range onlyA {
			if i > 0 {
				writer.WriteString(" ") // чтобы не ставить пробел перед первым элементом и
				// не ставить пробел после последнего элемента
			}
			writer.WriteString(strconv.Itoa(val))
		}
		writer.WriteString("\n")
	} else {
		writer.WriteString("0\n") // ← здесь печатаем "0" и переводим строку
	}

	// в печать уникальные b
	if len(onlyB) > 0 {
		writer.WriteString(strconv.Itoa(len(onlyB)) + "\n")
		for i, val := range onlyB {
			if i > 0 {
				writer.WriteString(" ") // чтобы не ставить пробел перед первым элементом и
				// не ставить пробел после последнего элемента
			}
			writer.WriteString(strconv.Itoa(val))
		}
		writer.WriteString("\n")
	} else {
		writer.WriteString("0\n") // ← здесь печатаем "0" и переводим строку
	}

}
