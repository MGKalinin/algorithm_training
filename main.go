package main

import (
	"bufio"
	"log"
	"os"
)

// 141. Правильная скобочная последовательность
func main() {
	// *строку--
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
