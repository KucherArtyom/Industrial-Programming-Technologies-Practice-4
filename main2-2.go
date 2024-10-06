package main

import (
	"fmt"
	"os"
)

func main() {
	var year int
	fmt.Println("Введите год: ")
	fmt.Fscan(os.Stdin, &year)

	if year%400 == 0 {
		fmt.Println("Данный год високосный!")
	} else if year%100 == 0 {
		fmt.Println("Данный год не является високосным!")
	} else if year%4 == 0 {
		fmt.Println("Данный год високосный!")
	} else {
		fmt.Println("Данный год не является високосным!")
	}
}
