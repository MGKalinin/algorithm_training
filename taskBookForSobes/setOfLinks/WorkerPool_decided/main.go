package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Интерактивная задача на кодинг, приближенный к условиям
// работы

// worker функция выполняет запросы
func worker(input <-chan string, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range input {
		resp, err := http.Get(url)
		if err != nil {
			output <- fmt.Sprintf("адрес %v - not ok\n", url)
			continue //Worker должен продолжать работу после обработки каждого URL
			// return был ошибкой - он завершал весь worker
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			output <- fmt.Sprintf("адрес %v - ok\n", url)
		} else {
			output <- fmt.Sprintf("адрес %v - not ok \n", url)
		}
	}
}

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

	const work = 5
	input := make(chan string, len(urls))
	output := make(chan string, len(urls))
	wg := sync.WaitGroup{}

	for i := 1; i <= work; i++ {
		wg.Add(1)
		go worker(input, output, &wg)
	}

	go func() {
		for _, url := range urls {
			input <- url
		}
		close(input)
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	for ans := range output {
		fmt.Print(ans)
	}
}
