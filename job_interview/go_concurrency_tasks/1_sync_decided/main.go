package main

import (
	"fmt"
	"sync"
)

// print square of range 0...20 in random order
//func main() {
//	counter := 20
//	var wg sync.WaitGroup
//	for i := 0; i < counter; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			fmt.Println(i * i)
//
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	counter := 20
	result := make(chan int, counter)
	var wg sync.WaitGroup

	// 1. Сначала запускаем ВСЕ рабочие горутины
	for i := 0; i < counter; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			result <- x * x
		}(i)
	}

	// 2. Затем горутина для закрытия канала
	go func() {
		wg.Wait()     // Ждём завершения ВСЕХ рабочих
		close(result) // Теперь канал закрывается безопасно
	}()

	// 3. Чтение данных (блокируется до появления данных)
	for v := range result {
		fmt.Println(v)
	}
}
