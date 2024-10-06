package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Println("Введите любое целое число: ")
	fmt.Fscan(os.Stdin, &number)
	if number%2 == 1 {
		fmt.Println("Введённое число нечётное!")
	} else {
		fmt.Println("Введённое число чётное!")
	}
}
