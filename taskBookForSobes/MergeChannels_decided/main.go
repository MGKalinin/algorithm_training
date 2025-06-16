package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	// Запускаем горутины для отправки данных в каналы
	go func() {
		defer close(ch1)
		defer close(ch2)
		defer close(ch3)
		ch1 <- 1
		ch1 <- 2
		ch2 <- 3
		ch2 <- 4
		ch3 <- 5
		ch3 <- 6
	}()

	for item := range merge(ch1, ch2, ch3) {
		fmt.Printf("Получено значение: %d\n", item)
	}
	fmt.Println("ok")
}

// Написать код функции, которая делает merge N каналов. Весь входной поток
// перенаправляется в один канал.
func merge(cs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				output <- i
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
