package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Введите число для вычисления факториала:")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Факториал отрицательных чисел не существует.")
	} else {
		fmt.Printf("Факториал числа %d равен %d\n", number, factorial(number))
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
