package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Println("Введите любое целое число: ")
	fmt.Fscan(os.Stdin, &number)
	if number%5 == 0 && number%3 == 0 {
		fmt.Println("Данное число делится одновременно на 3 и 5")
	} else {
		fmt.Println("Данное число не делится одновременно на 3 и 5")
	}
}
