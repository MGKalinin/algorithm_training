package main

import "fmt"

func main() {
	for item := range merge(...){
		// ...
	}
	fmt.Println("ok")
}

// Написать код функции, которая делает merge N каналов. Весь входной поток
// перенаправляется в один канал.
func merge(cs ...<-chan int) <-chan int {

}
