package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var mas int = 0
	var n1 int

	var p int = 0
	fmt.Println("Введите количество значений в массиве: ")
	fmt.Fscan(os.Stdin, &n)
	var numbers1 []int
	var numbers2 []int
	for i := 1; i <= n; i++ {
		fmt.Println("Введите", mas+1, "значение в массиве: ")
		fmt.Fscan(os.Stdin, &n1)
		numbers1 = append(numbers1, n1)
		numbers2 = append(numbers2, 0)
		mas = mas + 1
	}
	for i := 1; i <= n; i++ {
		numbers2[mas-1] = numbers1[p]
		p = p + 1
		mas = mas - 1
	}
	fmt.Println("Перевёрнутый массив: ", numbers2)
}
