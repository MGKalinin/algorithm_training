package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// // стек
// type Stack struct {
// 	elements []int
// }

// // метод push -добавить в стек число
// func (s *Stack) Push(element int) {
// 	s.elements = append(s.elements, element)
// }

// 140. Стек с защитой от ошибок
func main() {
	// defer fmt.Println("one")
	// defer fmt.Println("two")

	// One()
	// forRange()
	gracefulShutdown()
	fmt.Println("exit")
}

func One() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// wg.Add(1000)

	for i := 0; i < 1000; i++ {
		wg.Add(1) // добавлять до запуска горутины
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func forRange() {
	// show with for range buffered
	bufferedChannel := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}
	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel {
		fmt.Println("buffered", v)
	}

	unbufferedChannel := make(chan int)
	go func() {
		for _, num := range numbers {
			unbufferedChannel <- num
		}
		close(unbufferedChannel)
	}()

	for value := range unbufferedChannel {
		fmt.Println("unbuffered", value)
	}
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("time`s up")
		return
	case sig := <-sigChan:
		fmt.Println("Stopped by signal:", sig)
		return
	}
}
