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
```
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
```
4) Объединение строк
```
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
```
5) Расчет расстояния между двумя точками
```
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
```
### 2. Задачи с условным оператором

1) Проверка на четность/нечетность
```
package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Println("Введите любое целое число: ")
	fmt.Fscan(os.Stdin, &number)
	if number%2 == 1 {
		fmt.Println("Введённое число нечётное!")
	} else {
		fmt.Println("Введённое число чётное!")
	}
}
```
2) Проверка высокосного года
```
package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Println("Введите любое целое число: ")
	fmt.Fscan(os.Stdin, &number)
	if number%2 == 1 {
		fmt.Println("Введённое число нечётное!")
	} else {
		fmt.Println("Введённое число чётное!")
	}
}
```
3) Определение наибольшего из трех чисел
```
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
```
4) Категория возраста
```
package main

import (
	"fmt"
	"os"
)

// Категории возраста
// Малыш: 0 - 12
// Подросток: 13 - 19
// Взрослый: 20 - 55
// Пожилой: 56 - 100

func main() {
	var age int
	fmt.Println("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)
	if age >= 0 && age <= 12 {
		fmt.Println("Малыш")
	}
	if age >= 13 && age <= 19 {
		fmt.Println("Подросток")
	}
	if age >= 20 && age <= 55 {
		fmt.Println("Взрослый")
	}
	if age >= 56 && age < 100 {
		fmt.Println("Пожилой")
	}
}
```
5) Проверка делимости на 3 и 5
```
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
```
### 3. Задачи на циклы

1) Факториал числа
```
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Введите число для вычисления факториала:")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Факториал отрицательных чисел не существует.")
	} else {
		fmt.Printf("Факториал числа %d равен %d\n", number, factorial(number))
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
```
2) Числа Фибоначчи
```
package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var number1 = 0
	var number2 = 1
	var number3 = 0
	fmt.Println("Введите любое целое число n: ")
	fmt.Fscan(os.Stdin, &n)
	for i := 1; i <= n; i++ {
		fmt.Println(number1)
		number3 = number1 + number2
		number1 = number2
		number2 = number3
	}
}
```
3) Реверс массива
```
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
```
4) Поиск простых чисел
```
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
```
5) Сумма чисел в массиве
```
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
```
