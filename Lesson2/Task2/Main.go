package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	var d float64
	fmt.Println("Введите площадь")
	fmt.Scanln(&s)
	d = 2 * math.Sqrt(s/math.Pi)
	fmt.Printf("Диаметр прямоугольника равна - %f\n", d)
	fmt.Printf("Длинна окружности равна - %f", math.Pi*d)
}
