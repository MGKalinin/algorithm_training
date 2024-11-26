package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// стек
type Stack struct {
	elements []int
}

/*
метод push n
Добавить в стек число n (значение n задается после команды). Программа должна вывести ok.
*/
func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
	fmt.Println("ok")
}

/*
метод pop
Удалить из стека последний элемент. Программа должна вывести его значение.
Если во входных данных встречается операция back или pop, и при этом стек пуст,
то программа должна вместо числового значения вывести строку error.
*/
func (s *Stack) Pop() {
	n := len(s.elements)
	if n == 0 {
		fmt.Println("error")
		return
	} else {
		ans := s.elements[n-1]
		s.elements = s.elements[:(n - 1)]
		fmt.Println(ans)
	}
}

/*
back
Программа должна вывести значение последнего элемента, не удаляя его из стека.
Если во входных данных встречается операция back или pop, и при этом стек пуст,
то программа должна вместо числового значения вывести строку error.
*/
func (s *Stack) Back() {
	n := len(s.elements)
	if n == 0 {
		fmt.Println("error")
		return
	} else {
		ans := s.elements[n-1]
		fmt.Println(ans)
	}
}

/*
size
Программа должна вывести количество элементов в стеке.
*/
func (s *Stack) Size() {
	n := len(s.elements)
	fmt.Println(n)
}

/*
clear
Программа должна очистить стек и вывести ok.
*/
func (s *Stack) Clear() {
	s.elements = nil
	fmt.Println("ok")
}

/*
exit
Программа должна вывести bye и завершить работу.
*/
func (s *Stack) Exit() {
	fmt.Println("bye")
}

// 140. Стек с защитой от ошибок
func main() {
	stack := Stack{}
	// * строку --
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		parts := strings.Fields(command)
		switch parts[0] {
		case "push":
			num, _ := strconv.Atoi(parts[1])
			stack.Push(num)
		case "pop":
			stack.Pop()
		case "back":
			stack.Back()
		case "size":
			stack.Size()
		case "clear":
			stack.Clear()
		case "exit":
			stack.Exit()
			return

		}
	}

}
