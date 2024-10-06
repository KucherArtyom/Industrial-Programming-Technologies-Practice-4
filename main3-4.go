package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Println("Введите любое целое число:")
	fmt.Fscan(os.Stdin, &number)

	if number > 1 {
		fmt.Println("Простые числа до", number, ":")
		for i := 2; i <= number; i++ {
			isPrime := true
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				fmt.Print(i, " ")
			}
		}
		fmt.Println()
	} else {
		fmt.Println("Простых чисел нет")
	}
}
