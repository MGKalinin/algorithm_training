package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

// TODO писать код в песочнице
// Озон Платформа
//
// Дата: февраль 2025
// Грейд: мидл
// #Озон
//
// Задачи присланы участником сообщества. Как прислать свою задачу читайте тут.
//
// 1 Скрининг
// Дано x, определить, является ли x степенью двойки.
// Решение: x & (x - 1) == 0 – если выполняется, значит x степень двойки.
// 2 Техническое собеседование
// Задача 1:
// Дан список URL, нужно синхронно пройтись по нему и:
// Вывести "ОК", если статус-код 200.
// Вывести "не ОК", если статус-код  не 200 или произошла ошибка.
//
// Задача 2:
// Та же проверка, но теперь асинхронно/в параллель, при этом ничего не должно потеряться.
func main() {
	urls := []string{
		"https://yandex.ru",
		"https://habr.com",
		"https://lenta.ru",
		"https://ria.ru",
		"https://tass.ru",
		"https://ozone.ru",
	}

	// TODO 1.semaphore
	//channel for results of request
	//result := make(chan string, len(urls))
	//
	////semafore channel
	//semaf := make(chan struct{}, 3)
	//wg := sync.WaitGroup{}
	//
	//for _, url := range urls {
	//	wg.Add(1)
	//	go func(u string) {
	//		defer wg.Done()
	//		semaf <- struct{}{}
	//		defer func() { <-semaf }()
	//		//request by url
	//		req, err := http.Get(u)
	//		if err != nil {
	//			result <- fmt.Sprintf("%s: не ook", u)
	//			return
	//		}
	//		defer req.Body.Close()
	//		if req.StatusCode == http.StatusOK {
	//			result <- fmt.Sprintf("%s: ok", u)
	//		} else {
	//			result <- fmt.Sprintf("%s: не ОК", u)
	//		}
	//
	//	}(url)
	//}
	//go func() {
	//	wg.Wait()
	//	close(result)
	//}()
	//for val := range result {
	//	fmt.Println(val)
	//}

	// TODO 2.worker pool
	// https://go.dev/play/p/rbFgFYMKk6m
	//	const workers = 3
	//	n := len(urls)
	//
	//	//channel for tasks
	//	tasks := make(chan string, n)
	//	//channel for results
	//	result := make(chan string, n)
	//
	//	wg := sync.WaitGroup{}
	//
	//	//run the workers
	//	for i := 1; i < workers; i++ {
	//		wg.Add(1)
	//		go worker(tasks, result, &wg)
	//	}
	//
	//	//run cicle by slice and recive the answer
	//	for _, url := range urls {
	//		tasks <- url
	//	}
	//	close(tasks)
	//
	//	go func() {
	//		wg.Wait()
	//		close(result)
	//	}()
	//
	//	//read the answer
	//	for ans := range result {
	//		fmt.Println(ans)
	//	}
	//}
	//
	//// worker for requests
	//func worker(tasks <-chan string, result chan<- string, wg *sync.WaitGroup) {
	//	defer wg.Done()
	//
	//	for url := range tasks {
	//		req, err := http.Get(url)
	//		if err != nil {
	//			result <- fmt.Sprintf("not ook")
	//			continue
	//		}
	//
	//		defer req.Body.Close()
	//		if req.StatusCode == http.StatusOK {
	//			result <- fmt.Sprintf("ok")
	//		} else {
	//			result <- fmt.Sprintf("not ok")
	//		}
	//	}

	// TODO 3.использовать semaphore и контекст отмены после двух 200 ок
	// TODO to finalize code with context
	// https://go.dev/play/p/g420cG8vrVL

	//context for initialisation cancel with contex
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// var for count of StatuOK
	countOk := 0
	//add mutex for atomit operation with var of count
	var mu sync.Mutex

	// channel for semaphore len 3
	semaphore := make(chan struct{}, 3)
	//channel for result
	result := make(chan string, len(urls))

	//wg WaitGroup
	wg := sync.WaitGroup{}
	//run by slice
	wg.Add(len(urls))
	for _, url := range urls {
		go func(c string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			//check context situation
			select {
			case <-ctx.Done():
				fmt.Println("cancel situation")
				return
			default:
			}

			//make requests process
			request, err := http.NewRequestWithContext(ctx, "GET", c, nil)
			if err != nil {
				result <- fmt.Sprintf("not ook dont`t make request")
				return
			}
			req, err := http.DefaultClient.Do(request)
			if err != nil {
				result <- fmt.Sprintf("not oook mistake answer")
				return
			}
			defer req.Body.Close()
			//check the answer & incrementation of var of count
			if req.StatusCode == http.StatusOK {
				result <- fmt.Sprintf("ok")
				mu.Lock()
				countOk++
				if countOk >= 2 {
					cancel()
				}
				mu.Unlock()
			} else {
				result <- fmt.Sprintf("not ok")
			}

		}(url)
	}
	// close wg & result
	go func() {
		wg.Wait()
		close(result)
	}()
	//read results from result
	for ans := range result {
		fmt.Println(ans)
	}
}
