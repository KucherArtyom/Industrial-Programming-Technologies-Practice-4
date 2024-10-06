package main

import (
	"fmt"
	"os"
)

func fahrenheit(celsius_number float32) float32 {
	var fahrenheit_number = (1.8 * celsius_number) + 32
	return fahrenheit_number
}
func celsius(fahrenheit_number float32) float32 {
	var celsius_number = (0.5556 * (fahrenheit_number - 32))
	return celsius_number
}

func main() {
	var celsius_grad float32
	var fahrenheit_grad float32
	fmt.Println("Введите число градусов по Цельсию: ")
	fmt.Fscan(os.Stdin, &celsius_grad)
	fmt.Println(celsius_grad, "градусов по Цельсию, это ", fahrenheit(celsius_grad), "градусов по Фаренгейту")
	fmt.Println("Введите число градусов по Фаренгейту: ")
	fmt.Fscan(os.Stdin, &fahrenheit_grad)
	fmt.Println(fahrenheit_grad, "градусов по Фаренгейту, это ", celsius(fahrenheit_grad), "градусов по Цельсию")
}
