package main

import "fmt"

type Stack struct {
	elements []rune
}

// метод добавляет элемент в стек
func (s *Stack) Push(element rune) {
	s.elements = append(s.elements, element)
}

// 20. Valid Parentheses
func main() {
	s := "()"
	isValid(s)
}

// функ.проверки валидности
func isValid(s string) bool {
	stack := Stack{}
	openingBrackets := []rune{'(', '{', '['}
	closingBrackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	// если символ входящей строки равен rune из
	// openingBrackets - добавить в stack
	// для этого написать функцию -аргументы
	// слпайс rune и rune- вернуть true,
	//  если rune есть в слайс rune

	for _, char := range s {
		if checkingOccurrenceRune(openingBrackets, char) {
			stack.Push(char)
		}
	}
	fmt.Println(stack)
	return true
}

// функция проверяет вхождение rune в слайс rune
func checkingOccurrenceRune(arr []rune, r rune) bool {
	for _, val := range arr {
		if val == r {
			return true
		}
	}
	return false
}
