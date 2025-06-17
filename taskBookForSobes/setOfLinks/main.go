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

}
