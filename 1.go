package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Ошибка: неверный оператор")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f", a, operator, b, result)
}
