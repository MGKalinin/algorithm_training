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
	defer x.S()
	x.V = 456
}

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
//	callback := func() { data.Display() }
//	defer callback()
//	data.Value = 456
//	defer callback()
//}
