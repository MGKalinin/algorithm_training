package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 61. Пересечение множеств
func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	// первый слайс
	line1, _ := in.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	numbers1 := strings.Fields(line1)
	var slice1 []int
	for _, numStr := range numbers1 {
		nums, _ := strconv.Atoi(numStr)
		slice1 = append(slice1, nums)
	}
	// второй слайс
	line2, _ := in.ReadString('\n')
	line2 = strings.TrimSpace(line2)
	numbers2 := strings.Fields(line2)
	var slice2 []int
	for _, numStr := range numbers2 {
		nums, _ := strconv.Atoi(numStr)
		slice2 = append(slice2, nums)
	}
	sort.Ints(slice1)
	sort.Ints(slice2)
	//fmt.Println(slice1, slice2)

	first, second := 0, 0
	n := max(len(slice1), len(slice2))
	result := make([]int, 0, n)

	for first < len(slice1) && second < len(slice2) {
		if slice1[first] == slice2[second] {
			result = append(result, slice1[first])
			first++
			second++
		} else if slice1[first] < slice2[second] {
			first++
		} else {
			second++
		}
	}
	//fmt.Println(result)
	// Вывод результата согласно формату задачи
	if len(result) > 0 {
		for i, val := range result {
			if i > 0 {
				out.WriteString(" ")
			}
			out.WriteString(strconv.Itoa(val))
		}
		out.WriteString("\n")
	} else {
		out.WriteString("\n") // если пересечение пустое
	}
}
