package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	var sum int
	fmt.Println("Введите любое четырёхзначное, целое число: ")
	fmt.Fscan(os.Stdin, &number)
	sum = sum + (number % 10)
	number = number / 10
	sum = sum + (number % 10)
	number = number / 10
	sum = sum + (number % 10)
	number = number / 10
	sum = sum + (number % 10)
	number = number / 10
	fmt.Println("Сумма чисел введённого четырёхзначного, целого числа равна: ", sum)
}
