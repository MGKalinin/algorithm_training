package main

import "fmt"

// LeetCode 75
// Стеки / Stack

// 2390. Removing Stars From a String
func main() {
	arr := "leet**cod*e"
	removeStars(arr)
}

type Stack struct {
	elements []string
}

func (s *Stack) Push(elements []string, elem string) {
	elements = append(elements, elem)
}

func removeStars(s string) string {
	stack := Stack{}
	for _, char := range s {
		stack.Push(stack.elements, char)
	}
	fmt.Println(stack)
	if len(stack.elements) != 0 {
		return true
	}
	return false
}
