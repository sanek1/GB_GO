package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b float32
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, !): ")
	fmt.Scanln(&op)

	var res = calc(a, b, op)
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func calc(a, b float32, op string) float32 {
	var res float32
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("деление на 0 запрещено")
			os.Exit(1)
		}
		res = a / b
	case "!":
		res = float32(factorial(int(a)))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	return res
}

func factorial(n int) int {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
