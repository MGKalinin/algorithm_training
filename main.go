package main

import (
	"fmt"
	"strings"
)

func isCorrect(s string) string {
	// Создаем стек для хранения открывающих скобок
	stack := []rune{}

	// Создаем карту соответствия открывающих и закрывающих скобок
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// Проходим по каждому символу в строке
	for _, char := range s {
		// Если символ - открывающая скобка
		if _, ok := brackets[char]; ok {
			// Добавляем её в стек
			stack = append(stack, char)
		} else { // Если символ - закрывающая скобка
			// Проверяем, пустой ли стек
			if len(stack) == 0 {
				// Если стек пустой, значит, закрывающая скобка лишняя
				return "no"
			}

			// Извлекаем последнюю открывающую скобку из стека
			lastBracket := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Проверяем, соответствует ли закрывающая скобка открывающей
			if brackets[lastBracket] != char {
				// Если не соответствует, значит, последовательность неправильная
				return "no"
			}
		}
	}

	// После прохода по всей строке, стек должен быть пустым,
	// если все открывающие скобки были корректно закрыты
	if len(stack) == 0 {
		return "yes"
	} else {
		return "no"
	}
}

func main() {
	var input string
	fmt.Scanln(&input)
	input = strings.ReplaceAll(input, " ", "") //Удаляем пробелы если есть

	fmt.Println(isCorrect(input))
}
