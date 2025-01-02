package main

import (
	"fmt"
)

// C. Каждому по компьютеру https://contest.yandex.ru/contest/29075/problems/C/

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// // считать размер массива
	// line, _ := reader.ReadString('\n')
	// parts := strings.Fields(line)  //получается слайс строк без пробелов и символов новой строки
	// N, _ := strconv.Atoi(parts[0]) //N групп
	// M, _ := strconv.Atoi(parts[1]) //M аудиторий
	// // чтение массива
	// line, _ = reader.ReadString('\n')
	// parts = strings.Fields(line)
	// groupN := make([]int, N) //слайс челов по группам
	// for i := 0; i < N; i++ {
	// 	groupN[i], _ = strconv.Atoi(parts[i])
	// }
	// //слайс компьютеров в аудиториях
	// line, _ = reader.ReadString('\n')
	// parts = strings.Fields(line)
	// roomsM := make([]int, M)
	// for i := 0; i < M; i++ {
	// 	roomsM[i], _ = strconv.Atoi(parts[i])
	// }

	// Выведите на первой строке число P - количество групп, которые удастся распределить по аудиториям
	// На второй строке выведите распределение групп по аудиториям – N чисел,
	// i-е число должно соответствовать номеру аудитории, в которой должна заниматься i-я группа.
	// Если i-я группа осталась без аудитории, i-е число должно быть равно 0
	// M>=N

	N := 3
	M := 3
	groupN := []int{1, 2, 3} // слайс  групп
	roomsM := []int{3, 4, 2} // слайс аудиторий

	//возможность разместить
	abilityPlace := func(roomsSize, groupSize int) bool {
		return roomsSize >= groupSize+1
	}

	P := 0
	ans := make([]int, N)
	for indGroup := 0; indGroup < N; indGroup++ {
		fmt.Println("indGroup", indGroup)
		var foundGroup bool
		for indRooms := 0; indRooms < M; indRooms++ {
			fmt.Println("indRooms", indRooms)
			if abilityPlace(roomsM[indRooms], groupN[indGroup]) {
				ans[indGroup] = indRooms + 1
				P++
				foundGroup = true
				// добавить поиск наименьшего значения
				// аудитории

				// заменить значение аудитории
				// где размещена группа  на 0
				roomsM[indRooms] = 0
				break
			}
		}
		if !foundGroup {
			ans[indGroup] = 0
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

// тест 5
// 3 3
// 1 2 3
// 3 4 2
