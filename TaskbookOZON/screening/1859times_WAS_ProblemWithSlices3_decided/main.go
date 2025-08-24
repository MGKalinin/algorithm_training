package main

import "time"

// Что выведет программа? Почему?
func main() {
	timeStart := time.Now()
	_, _ = <-worker(), <-worker()
	println(int(time.Since(timeStart).Seconds())) // 6 сек
}
func worker() chan int {
	ch := make(chan int) // не буф канал
	go func() {
		time.Sleep(3 * time.Second) //3 сек 1 раз, потом 2 раз 3 сек
		ch <- 1
	}()
	return ch
}
