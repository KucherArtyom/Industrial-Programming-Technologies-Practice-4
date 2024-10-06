package main

import (
	"fmt"
	"os"
)

func main() {
	var lines [3]string
	fmt.Println("Введите первую строку: ")
	fmt.Fscan(os.Stdin, &lines[0])
	fmt.Println("Введите вторую строку: ")
	fmt.Fscan(os.Stdin, &lines[1])
	fmt.Println("Введите третью строку: ")
	fmt.Fscan(os.Stdin, &lines[2])
	fmt.Println(lines[0], lines[1], lines[2])
}
