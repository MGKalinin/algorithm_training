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

	// паттерн Semaphore
	var wg sync.WaitGroup
	const quantity = 3
	sem := make(chan struct{}, quantity)

	result := make(chan string, len(urls))
	for _, url := range urls {
		wg.Add(1)             //увеличивает счетчик ожидаемых горутин
		go func(url string) { //Все 8 горутин запускаются сразу в цикле(8 адресов)
			defer wg.Done()

			sem <- struct{}{} // Когда семафор заполнен (3 элемента),
			// операция sem <- блокирует выполнение горутины
			defer func() { <-sem }() //При освобождении слота (<-sem) одна из заблокированных горутин "просыпается"

			resp, err := http.Get(url)
			if err != nil {
				result <- fmt.Sprintf("адрес %v - not ok\n", url)
				return // Завершаем ЭТУ горутину (она обработала свой URL)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				result <- fmt.Sprintf("адрес %v - ok\n", url)
			} else {
				result <- fmt.Sprintf("адрес %v - not ok \n", url)
			}
		}(url)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	for ans := range result {
		fmt.Println(ans)
	}
}
