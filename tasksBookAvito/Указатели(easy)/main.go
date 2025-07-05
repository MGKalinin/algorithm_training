package main

import "fmt"

// Что выведет следующая программа и почему:

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
}
func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}

//Как модифицировать программу, чтобы вывелось:
//Bob
//Alice
