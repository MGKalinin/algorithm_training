package main

import (
	"fmt"
	"sync"
)

// ===========================================================
// Задача 1
// 1. Как это работает, что не так, что поправить?
// ===========================================================
func main() {
	//обзор проблем:

	//ch := make(chan bool)
	//ch <- true //1.после этого будет deadlock ,т.к.нужно будет прочитать из не буф.канала после записи
	//go func() {
	//	<-ch
	//}() //2.нет ожидания завершения горутины-она прервётся как завершится main
	////3.go func() может не запуститься до ch <- true.
	////В Go нет упорядочивания между запуском горутины и другими операциями
	//ch <- true //4.после этого тоже будет deadlock - если будет только одно чтение
	//	5.нет закрытия канала

	//	исправления:

	ch := make(chan bool)
	wg := sync.WaitGroup{}
	//ch <- true//1.перенесена после горутины чтения
	wg.Add(1) //2.добавить 1 горутину
	go func() {
		defer wg.Done() //3.отложенное закрытие
		for range ch {  //4.чтение из канала
			fmt.Println(ch)
		}
	}()
	ch <- true //1.вот сюда
	ch <- true
	close(ch) //5.закрыть канал
	wg.Wait() //6.дождаться завершения горутины
}

//
//===========================================================
//Задача 2
//1. Как будет работать код?
//2. Как сделать так, чтобы выводился только первый ch?
//===========================================================
//
//func main() {
//ch := make(chan bool)
//ch2 := make(chan bool)
//ch3 := make(chan bool)
//go func() {
//ch <- true
//}()
//go func() {
//ch2 <- true
//}()
//go func() {
//ch3 <- true
//}()
//
//select {
//case <-ch:
//fmt.Printf("val from ch")
//case <-ch2:
//fmt.Printf("val from ch2")
//case <-ch3:
//fmt.Printf("val from ch3")
//}
//}
//
//
//===========================================================
//Задача 3
//1. Что выведет код?
//===========================================================
//
//func main() {
//var m map[string]int
//for _, word := range []string{"hello", "world", "from", "the",
//"best", "language", "in", "the", "world"} {
//m[word]++
//}
//for k, v := range m {
//fmt.Println(k, v)
//}
//}
//
//===========================================================
//Задача 4
//1. Что выведет код и как исправить?
//===========================================================
//
//var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
//var a = 0
//
//func main() {
//wg := sync.WaitGroup{}
//wg.Add(3)
//go func() {
//wg.Done()
//a=10
//globalMap["test"] = append(globalMap["test"], a)
//
//}()
//go func() {
//wg.Done()
//a=11
//globalMap["test2"] = append(globalMap["test2"], a)
//}()
//go func() {
//wg.Done()
//a=12
//globalMap["test3"] = append(globalMap["test3"], a)
//}()
//wg.Wait()
//fmt.Printf("%v", globalMap)
//fmt.Printf("%d", a)
//}
//
//===========================================================
//Задача 5
//===========================================================
//
//type Result struct{}
//
//type SearchFunc func(ctx context.Context, query string) (Result, error)
//
//func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
//// Нужно реализовать функцию, которая выполняет поиск query во всех переданных SearchFunc
//// Когда получаем первый успешный результат - отдаем его сразу. Если все SearchFunc отработали
//// с ошибкой - отдаем последнюю полученную ошибку
//}
