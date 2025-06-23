package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

// Есть набор урлов.
func main() {
	//var urls = []string{
	//	"http://ozon.ru",
	//	"https://ozon.ru",
	//	"http://google.com",
	//	"http://somesite.com",
	//	"http://non-existent.domain.tld",
	//	"https://ya.ru",
	//	"http://ya.ru",
	//	"http://ёёёё",
	//}
	//3. Модифицируйте программу таким образом, чтобы нигде не использовалась длина
	//слайса урлов. Считайте, что урлы приходят из внешнего источника. Сколько их будет
	//заранее - неизвестно. Предложите идиоматичный вариант, как ваша программа будет
	//узнавать об окончании списка и передавать сигнал об окончании действий далее.
	//4. (необязательно, можно обсудить устно, чтобы убедиться, что кандидат понимает идею
	//контекста, либо предложить как домашнее задание) Модифицируйте программу таким
	//образом, что бы при получении 2 первых ответов с "200 OK" остальные запросы штатно
	//прерывались.
	//	При этом необходимо напечатать на экране сообщение о завершении запроса.

	//18 грейд
	//Написал синхронное решение, которое идет последовательно и выполняет http-запросы.
	//Написал решение с использованием канала (куда отправляются ссылки для скачивания),
	//горутин (воркеров), которые выполняют запросы и sync.WaitGroup для ожидания
	//завершения.
	//	Может реализовать остановку выполнения воркеров через context или управляющий
	//канал.

	result := inputData(os.Stdin)
	fmt.Println(result)
	//	2.реализовать worker pool: запустить заданное количество воркеров
	const work = 5
	input := make(chan string)
	output := make(chan string)
	wg := sync.WaitGroup{}

	for i := 1; i <= work; i++ {
		wg.Add(1)
		go worker(input, output, &wg)
	}

	go func() {
		for _, url := range result {
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
	//3.при получении 2 первых ответов с "200 OK" остальные запросы штатно
	//прерывались; реализовать остановку выполнения воркеров через context или управляющий канал.
	//TODO реализовать пункт 3
}

// 1.написать функцию считывающую вводимые данные
func inputData(r io.Reader) []string {
	urls := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
		if url == "" {
			fmt.Printf("Ввод окончен")
			break
		}
	}
	//fmt.Println(urls)
	return urls
}

// 2a worker функция выполняет запросы
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
