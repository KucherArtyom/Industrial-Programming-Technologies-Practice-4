package main

import (
	"fmt"
	"os"
)

func main() {
	var number1 float32
	var number2 float32
	var number3 float32
	fmt.Println("Введите первый номер: ")
	fmt.Fscan(os.Stdin, &number1)
	fmt.Println("Введите второй номер: ")
	fmt.Fscan(os.Stdin, &number2)
	fmt.Println("Введите третий номер: ")
	fmt.Fscan(os.Stdin, &number3)
	if number1 > number2 && number1 > number3 {
		fmt.Println(number1)
	} else if number2 > number3 {
		fmt.Println(number2)
	} else {
		fmt.Println(number3)
	}
}
