package main

import (
	"context"
	"fmt"
	"sync"
)

// 1. Смерджить 2 упорядоченных массива в 1 упорядоченный массив.
// Пример:
// дано: [1, 2, 3, 4, 5], [4, 5, 6, 7,8]
// Результат: [1, 2, 3, 4, 4, 5, 5, 6, 7, 8]
// TODO решить в песочнице
func merge(a []int, b []int) []int {

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
