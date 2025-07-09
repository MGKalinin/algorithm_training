package main

import (
	"fmt"
	"sync"
)

// Написать код функции, которая делает merge N каналов. Весь входной поток
// перенаправляется в один канал.
func merge(cs ...<-chan int) <-chan int {
	result := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	for _, ch := range cs {
		go func(c <-chan int) {
			for val := range c {
				result <- val
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := merge(ch1, ch2)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)

	}()
	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	for val := range merged {
		fmt.Println(val)
	}
}
