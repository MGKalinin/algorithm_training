package main

import "fmt"

// Что такое defer? Что выведет следующий код?

type X struct {
	V int
}

func (x X) S() {
	fmt.Println(x.V)
}

func main() {
	x := X{123}
	defer x.S() //выведет 123
	x.V = 456   //456 не будет выведено в печать т.к. присваивается значение 456,
	// но не вызывается метод x.S() выводящий в печать
}

//выведет 123

//на собесе было так:

//type Data struct {
//	Value int
//}
//
//func (x Data) Display() {
//	fmt.Println(x.Value)
//}
//
//func main() {
//	data := Data{123}
//	callback := func() { data.Display() } // Замыкание захватывает `data`.
//	// Если замыкание передаётся в defer, оно не копирует переменные, а хранит ссылку на них.
//	defer callback()
//	data.Value = 456 // (1) Будет выполнено последним
//	defer callback() // (2) Будет выполнено первым
//}

// Выведет 456 456 : callback ссылается на data, а не на её копию.
//В момент выполнения defer (при выходе из main) data.Value уже равно 456, поэтому оба вызова печатают 456.
