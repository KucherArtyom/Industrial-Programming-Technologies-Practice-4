package main

import (
	"fmt"
	"os"
)

// Категории возраста
// Малыш: 0 - 12
// Подросток: 13 - 19
// Взрослый: 20 - 55
// Пожилой: 56 - 100

func main() {
	var age int
	fmt.Println("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	if age >= 0 && age <= 12 {
		fmt.Println("Малыш")
	}
	if age >= 13 && age <= 19 {
		fmt.Println("Подросток")
	}
	if age >= 20 && age <= 55 {
		fmt.Println("Взрослый")
	}
	if age >= 56 && age < 100 {
		fmt.Println("Пожилой")
	}
}
