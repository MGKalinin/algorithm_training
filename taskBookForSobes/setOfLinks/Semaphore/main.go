package main

import (
	"fmt"
	"net/http"
	"sync"
)

//Интерактивная задача на кодинг, приближенный к условиям
//работы

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	//Напишите программу, которая:
	//2. Модифицируйте программу таким образом, чтобы использовались каналы для
	//коммуникации основного потока с горутинами. Пример:
	//•
	//Запросы по списку выполняются в горутинах.
	//•
	//Печать результатов на экран происходит в основном потоке

	// Semaphore
	var wg sync.WaitGroup
	const quantity = 3
	sem := make(chan struct{}, quantity)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("адрес %v - not ok\n", url)
			}
			resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				fmt.Printf("адрес %v - ok\n", url)
			} else {
				fmt.Printf("адрес %v - not ok \n", url)
			}
		}(url)
		go func() {
			wg.Wait()
		}()
	}
}
