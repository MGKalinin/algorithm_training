package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// C. Каждому по компьютеру https://contest.yandex.ru/contest/29075/problems/C/

func main() {
	reader := bufio.NewReader(os.Stdin)
	// считать размер массива
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)  //получается слайс строк без пробелов и символов новой строки
	N, _ := strconv.Atoi(parts[0]) //N групп
	M, _ := strconv.Atoi(parts[1]) //M аудиторий
	// чтение массива
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	groupN := make([]int, N) //слайс челов по группам
	for i := 0; i < N; i++ {
		groupN[i], _ = strconv.Atoi(parts[i])
	}
	//слайс компьютеров в аудиториях
	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	roomsM := make([]int, M)
	for i := 0; i < M; i++ {
		roomsM[i], _ = strconv.Atoi(parts[i])
	}

	// Выведите на первой строке число P - количество групп, которые удастся распределить по аудиториям
	// На второй строке выведите распределение групп по аудиториям – N чисел,
	// i-е число должно соответствовать номеру аудитории, в которой должна заниматься i-я группа.
	// Если i-я группа осталась без аудитории, i-е число должно быть равно 0
	// M>=N

	// N := 10
	// M := 10
	// groupN := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1} // слайс  групп
	// roomsM := []int{2, 2, 2, 2, 2, 1, 2, 2, 2, 2} // слайс аудиторий
	P := 0
	ans := []int{}
	for i := 0; i < M; i++ {
		if groupN[i] < roomsM[i] { //если группа меньше компьютеров в аудитории добавить в слай ответов номер аудитории
			ans = append(ans, i+1)
			P++
		} else if groupN[i] >= roomsM[i] { //группа равна либо больше компьютеров в аудитории добавить 0
			ans = append(ans, 0)
		}
	}

	fmt.Println(P)
	// fmt.Println(ans)
	for _, val := range ans {
		fmt.Println(val)
	}
}

// тест 4
// 10 10 N групп, M аудиторий
// 1 1 1 1 1 1 1 1 1 1 человек в каждой группе
// 2 2 2 2 2 2 2 2 2 1 компьютеров в каждой аудитории
// какое кол-во групп разместить? +один комп для учителя
