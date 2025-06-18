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
	//1. Поочередно выполнит http запросы по предложенному списку ссылок
	//•
	//в случае получения http-кода ответа на запрос "200 OK" печатаем на экране "адрес url -
	//ok"
	//•
	//в случае получения http-кода ответа на запрос отличного от "200 OK" либо в случае
	//ошибки печатаем на экране "адрес url - not ok"

	//for _, url := range urls {
	//	response, err := http.Get(url)
	//	if err != nil {
	//		fmt.Printf("адрес  %v - not ok\n", url)
	//		continue
	//	}
	//	defer response.Body.Close()
	//	// Добавляем проверку статус-кода!
	//	if response.StatusCode == http.StatusOK { // 200
	//		fmt.Printf("адрес %v - ok\n", url)
	//	} else {
	//		fmt.Printf("адрес %v - not ok (статус: %d)\n", url, response.StatusCode)
	//	}
	//
	//}

	//2. Модифицируйте программу таким образом, чтобы использовались каналы для
	//коммуникации основного потока с горутинами. Пример:
	//•
	//Запросы по списку выполняются в горутинах.
	//•
	//Печать результатов на экран происходит в основном потоке

	//	используя Worker Pool ==============================
	//var wg sync.WaitGroup
	////Определяем количество воркеров (например, 3-5)
	//const numberWorker = 5
	////Канал для задач (tasks): chan string
	//tasksChan := make(chan string)
	////Канал для результатов (results): chan string
	//resultsChan := make(chan string)
	////Создаем N горутин-воркеров
	//for i := 0; i < numberWorker; i++ {
	//	wg.Add(1)
	//	go worker(tasksChan, resultsChan, &wg)
	//}
	//// В отдельной горутине: Последовательно отправляет все URL в канал задач
	//go func() {
	//	for _, url := range urls {
	//		tasksChan <- url
	//	}
	//	close(tasksChan) // закрыть канал задач
	//}()
	//// Ожидания завершения всех воркеров & закрыть канал результатов
	//go func() {
	//	wg.Wait()
	//	close(resultsChan)
	//}()
	//// Основной поток: Читает результаты из канала результатов
	//for answer := range resultsChan {
	//	fmt.Println(answer)
	//}

	//используя Semaphore ==============================
	const maxParallel = 3 // Максимальное количество одновременных запросов
	//семафор
	sem := make(chan struct{}, maxParallel)
	//Создать канал для результатов (буферизованный)
	resultsChan := make(chan string)
	//	Создать sync.WaitGroup для отслеживания завершения
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			// Захватываем слот в семафоре (ограничиваем параллелизм)
			sem <- struct{}{}
			// Гарантируем освобождение слота при выходе из функции
			defer func() { <-sem }()

			// Выполняем HTTP-запрос
			response, err := http.Get(u)
			// Обработка ошибок сети/DNS
			if err != nil {
				// Если есть response (например, при ошибке редиректа), закрываем тело
				if response != nil {
					response.Body.Close()
				}
				resultsChan <- fmt.Sprintf("адрес %s - not ok (ошибка: %v)", u, err)
				return
			}
			// Гарантированно закрываем тело ответа
			defer response.Body.Close()
			// Проверяем статус код
			if response.StatusCode == http.StatusOK {
				resultsChan <- fmt.Sprintf("адрес %s - ok", u)
			} else {
				resultsChan <- fmt.Sprintf("адрес %s - not ok (статус: %d)", u, response.StatusCode)
			}

		}(url)
	}
	// Запускаем горутину для закрытия канала результатов после завершения всех запросов
	go func() {
		wg.Wait()          // Ждем завершения всех горутин
		close(resultsChan) // Закрываем канал результатов
	}()

	// Основной поток: читаем и выводим результаты
	for result := range resultsChan {
		fmt.Println(result)
	}
}

//// функция worker отправляет запрос, читает результат
//func worker(tasksChan <-chan string, resultsChan chan<- string, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for url := range tasksChan {
//		response, err := http.Get(url)
//		if err != nil {
//			resultsChan <- fmt.Sprintf("адрес %s - not ok (ошибка: %v)", url, err)
//			continue
//		}
//		response.Body.Close()
//		if response.StatusCode == http.StatusOK {
//			resultsChan <- fmt.Sprintf("адрес %s - ok", url)
//		} else {
//			resultsChan <- fmt.Sprintf("адрес %s - not ok (статус: %d)", url, response.StatusCode)
//		}
//	}
//}
