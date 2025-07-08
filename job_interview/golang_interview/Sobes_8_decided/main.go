package main

import "fmt"

// ===========================================================
// Задача 1_workerPool_semafore
// 1_workerPool_semafore. Добавить код, который выведет тип переменной whoami
// ===========================================================
//func printType(whoami interface{}) {
//	//fmt.Printf("%T\n", whoami)
//}

// ===========================================================
// Задача 2
// 1_workerPool_semafore. Что выведет код?
// ===========================================================
func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

//start
//end
//3
//2
//1_workerPool_semafore
