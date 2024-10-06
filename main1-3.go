package main

import (
	"fmt"
	"os"
)

func main() {
	var numbers [3]int
	fmt.Println("Введите первый элемент трёхмерного массива: ")
	fmt.Fscan(os.Stdin, &numbers[0])
	numbers[0] = numbers[0] * 2
	fmt.Println("Введите второй элемент трёхмерного массива: ")
	fmt.Fscan(os.Stdin, &numbers[1])
	numbers[1] = numbers[1] * 2
	fmt.Println("Введите третий элемент трёхмерного массива: ")
	fmt.Fscan(os.Stdin, &numbers[2])
	numbers[2] = numbers[2] * 2
	fmt.Println("Преобразованный массив: ", numbers)
}
