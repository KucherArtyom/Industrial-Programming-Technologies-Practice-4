package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	var hypotenuse float64
	fmt.Println("Введите координату x первой точки: ")
	fmt.Fscan(os.Stdin, &x1)
	fmt.Println("Введите координату y первой точки: ")
	fmt.Fscan(os.Stdin, &y1)
	fmt.Println("Введите координату x второй точки: ")
	fmt.Fscan(os.Stdin, &x2)
	fmt.Println("Введите координату y второй точки: ")
	fmt.Fscan(os.Stdin, &y2)
	var leg1 = x2 - x1
	var leg2 = y2 - y1
	hypotenuse = (leg1 * leg1) + (leg2 * leg2)
	fmt.Println("Расстояние от первой точки  до второй точки в 2D пространстве = ", (math.Sqrt(hypotenuse)))
}
