package main

import (
	"fmt"
)

// 1. Смерджить 2 упорядоченных массива в 1 упорядоченный массив.
// Пример:
// дано: [1, 2, 3, 4, 5], [4, 5, 6, 7,8]
// Результат: [1, 2, 3, 4, 4, 5, 5, 6, 7, 8]
// TO DO решить в песочнице https://go.dev/play/p/SR093pFtjX1

func merge(a []int, b []int) []int {
	n := len(a) + len(b)
	ans := make([]int, 0, n)

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i++
		} else {
			ans = append(ans, b[j])
			j++
		}
	}

	//for ; i < len(a); i++ {
	//ans = append(ans, a[i])
	//}

	//for ; j < len(b); j++ {
	//ans = append(ans, b[j])
	//}

	ans = append(ans, a[i:]...)
	//result = append(result, a[i:]...)
	//Эта конструкция делает следующее:
	//a[i:] - создает новый срез (slice), который содержит:
	//Все элементы исходного среза a
	//Начиная с индекса i
	//До конца среза
	//... - оператор распаковки (variadic operator). Он "распаковывает" элементы среза, передавая их как отдельные аргументы.
	//	append - добавляет распакованные элементы в конец среза result
	ans = append(ans, b[j:]...)

	fmt.Println(ans)
	return ans
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}
	fmt.Println(merge(a, b))
}

// 2. Смерджить произвольное количество каналов в один канал
// TO DO для задачи слияния произвольного количества каналов в один паттерн Fan-In идеально подходит
// написано в песочнице ) https://go.dev/play/p/UupcgiLal7k

//	func merge(chList ...chan int) chan int {
//		out := make(chan int, len(chList)) // это не правильно-каналов может быть два, а элементов по 10 в каждом
//		// прописывать буфер с запасом или предполагать что чтение из не буферизированного канала оперативно
//
//		out := make(chan int)
//		wg := sync.WaitGroup{}
//		for _, ch := range chList {
//			wg.Add(1)
//			go func(c chan int) {
//				defer wg.Done()
//				for val := range c {
//					out <- val
//				}
//			}(ch)
//		}
//		go func() {
//			wg.Wait()
//			close(out)
//		}()
//
//		return out
//	}
//
// 2а. TO DO выполнить ту же задачу с использованием контекста
// https://go.dev/play/p/bSroc24u9RK
//func merge(ctx context.Context, chList ...<-chan int) <-chan int {
//	out := make(chan int)
//	wg := sync.WaitGroup{}
//	for _, ch := range chList {
//		wg.Add(1)
//		go func(c <-chan int) {
//			defer wg.Done()
//			for {
//				select {
//				case val, ok := <-c:
//					{
//						if !ok {
//							return // channel is close
//						}
//						select {
//						case out <- val:
//						case <-ctx.Done():
//							return // context
//						}
//					}
//
//				case <-ctx.Done():
//					return // context
//				}
//
//			}
//		}(ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(out)
//	}()
//
//	return out
//}
//
//func main() {
//	ctx := context.Background()
//
//	a := make(chan int, 2)
//	a <- 1
//	a <- 2
//	close(a)
//
//	b := make(chan int, 1)
//	b <- 1
//	close(b)
//
//	res := merge(ctx, a, b)
//	for v := range res {
//		fmt.Println(v)
//	}
//}
