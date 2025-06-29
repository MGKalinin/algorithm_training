package main

import (
	"fmt"
	"net/http"
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

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("адрес %v - not ok\n", url)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("адрес %v - ok\n", url)
		} else {
			fmt.Printf("адрес %v - not ok \n", url)
		}
	}
}
