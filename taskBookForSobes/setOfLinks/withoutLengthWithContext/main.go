package main

import (
	"bufio"
	"context"
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
	count := 0
	//	2.реализовать worker pool: запустить заданное количество воркеров
	const work = 5
	input := make(chan string)
	output := make(chan Response)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= work; i++ {
		wg.Add(1)
		go worker(ctx, input, output, &wg)
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

	//3.при получении 2 первых ответов с "200 OK" остальные запросы штатно
	//прерывались; реализовать остановку выполнения воркеров через context или управляющий канал.
	//count := 0
	for resp := range output {
		fmt.Print(resp.Message)
		if resp.Status == http.StatusOK {
			count++
			fmt.Printf("Received 200 OK, count: %d\n", count)
			if count == 2 {
				cancel() // Отменяем контекст после получения двух 200 OK
			}
		}
	}
	fmt.Println("All workers stopped\n")
}

// 1.написать функцию считывающую вводимые данные
func inputData(r io.Reader) []string {
	urls := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
		if url == "" {
			fmt.Printf("Ввод окончен\n")
			break
		}
	}
	//fmt.Println(urls)
	return urls
}

type Response struct {
	URL     string
	Status  int
	Message string
}

// 2a worker функция выполняет запросы
func worker(ctx context.Context, input <-chan string, output chan<- Response, wg *sync.WaitGroup) {
	// Шаг 1: Уменьшаем счетчик WaitGroup при завершении
	defer wg.Done()

	// Шаг 2: Основной рабочий цикл
	for {
		select {
		// Шаг 3: Обработка получения URL из входного канала
		case url, ok := <-input:
			// Шаг 4: Проверка закрытия канала
			if !ok {
				return // Канал закрыт - завершаем работу
			}

			// Шаг 5: Создание HTTP-запроса с контекстом
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				// Шаг 6: Обработка ошибки создания запроса
				output <- Response{
					URL:     url,
					Status:  0,
					Message: fmt.Sprintf("адрес %v - ошибка создания запроса: %v\n", url, err),
				}
				continue
			}

			// Шаг 7: Выполнение HTTP-запроса
			resp, err := http.DefaultClient.Do(req)

			// Шаг 8: Обработка ошибок выполнения запроса
			if err != nil {
				message := fmt.Sprintf("адрес %v - ошибка выполнения: %v\n", url, err)
				// Проверяем отмену контекста
				if ctx.Err() == context.Canceled {
					message = fmt.Sprintf("адрес %v - запрос отменен\n", url)
				}

				output <- Response{
					URL:     url,
					Status:  0,
					Message: message,
				}
				continue
			}

			// Шаг 9: Гарантируем закрытие тела ответа
			defer resp.Body.Close()

			// Шаг 10: Формирование успешного ответа
			output <- Response{
				URL:     url,
				Status:  resp.StatusCode,
				Message: fmt.Sprintf("адрес %v - %s\n", url, http.StatusText(resp.StatusCode)),
			}

		// Шаг 11: Обработка отмены контекста
		case <-ctx.Done():
			return // Немедленное завершение при отмене контекста
		}
	}
}
