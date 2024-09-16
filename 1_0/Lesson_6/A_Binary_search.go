// A. Двоичный поиск
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	N, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
	}
	K, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(N, K)

	// список из N
	scanner.Scan()
	line = scanner.Text()
	parts = strings.Fields(line)
	arr_1 := make([]int, N)
	for i, part := range parts {
		arr_1[i], err = strconv.Atoi(part)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
	sort.Ints(arr_1)
	fmt.Println(arr_1)
	// список из K
	scanner.Scan()
	line = scanner.Text()
	parts = strings.Fields(line)
	arr_2 := make([]int, K)
	for i, part := range parts {
		arr_2[i], err = strconv.Atoi(part)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
	sort.Ints(arr_2)
	fmt.Println(arr_2)

}
