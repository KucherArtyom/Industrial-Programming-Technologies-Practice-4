# Industrial-Programming-Technologies-Practice-4 (ЭФМО-02-24 Кучер Артем Сергеевич)
## Задачи для практической работы на языке Go
### 1. Задачи на линейное программирование (без условных операторов и циклов)

1) Сумма цифр числа
```
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
```
2) Преобразование температуры
```
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
```
3) Удвоение каждого элемента массива
'''
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
'''
4) Объединение строк

5) Расчет расстояния между двумя точками

### 2. Задачи с условным оператором

1) Проверка на четность/нечетность

2) Проверка высокосного года

3) Определение наибольшего из трех чисел

4) Категория возраста

5) Проверка делимости на 3 и 5

### 3. Задачи на циклы

1) Факториал числа

2) Числа Фибоначчи

3) Реверс массива

4) Поиск простых чисел

5) Сумма чисел в массиве
