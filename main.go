package main

import "fmt"

// стек
type Stack struct {
	elements []int
}

/*
метод push n
Добавить в стек число n (значение n задается после команды). Программа должна вывести ok.
*/
func (s *Stack) Push(command string, element int) {
	s.elements = append(s.elements, element)
	fmt.Println("ok")
}

/*
метод pop
Удалить из стека последний элемент. Программа должна вывести его значение.
*/
func (s *Stack) Pop(command string) {
	n := len(s.elements)
	ans := s.elements[n-1]
	s.elements = s.elements[:(n - 1)]
	fmt.Println(ans)
}

// 140. Стек с защитой от ошибок
func main() {

}
