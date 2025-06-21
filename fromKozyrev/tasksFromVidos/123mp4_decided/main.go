package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.написать код для выдачи массива уникальных случайных чисел
//func numberGenerator(n int) ([]int, error) {
//	if n < 0 {
//		return nil, errors.New("errors: n must >0")
//	}
//	if n == 0 {
//		return []int{}, nil
//	}
//	unicNums := make(map[int]struct{})
//	slice := make([]int, 0, n)
//	for len(slice) < n {
//		num := rand.Int() // Случайное целое (int)
//		if _, ok := unicNums[num]; !ok {
//			unicNums[num] = struct{}{}
//			slice = append(slice, num)
//		}
//	}
//	return slice, nil
//}
//func main() {
//	rand.New(rand.NewSource(time.Now().UnixNano()))
//	n := 10
//	fmt.Println(numberGenerator(n))
//}

//3.реализовать worker pool. есть 10 задач(функций) каждая засыпает на 1 секунду и
//выводит номер воркера который эту задачу исполнил, количество воркеров задаётся при запуске

// worker функция выполняет задачу заснуть на секунду и вывести номер выполнившего задачу воркера
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for taskID := range tasks {
		time.Sleep(1 * time.Second)
		fmt.Printf("задачу %d выполнил воркер %d\n", taskID, id)
	}
}

func main() {
	//количество задач
	const tasksNum = 10
	//количество воркеров
	const workers = 3
	//tasks канал задач
	tasks := make(chan int, tasksNum)

	wg := sync.WaitGroup{}
	//	запускаем заданное количество воркеров
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}
	//	запуск заданного количества задач
	for i := 0; i < tasksNum; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()

}
