package main

import (
	"fmt"
	"os"
)

func main() {
	var sum int
	var mas_num int
	var mas int = 0
	var n1 int
	var mas2 int = 0
	fmt.Println("Введите количество значений в массиве: ")
	fmt.Fscan(os.Stdin, &mas_num)
	var numbers1 []int
	for i := 1; i <= mas_num; i++ {
		fmt.Println("Введите", mas+1, "значение в массиве: ")
		fmt.Fscan(os.Stdin, &n1)
		numbers1 = append(numbers1, n1)
		mas = mas + 1
	}
	for i := 1; i <= mas_num; i++ {
		sum = sum + numbers1[mas2]
		mas2 = mas2 + 1
	}
	fmt.Println("Сумма чисел в массиве: ", sum)
}
