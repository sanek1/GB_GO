package main

import (
	"fmt"
)

func main() {
	var width, height int
	fmt.Println("Введите ширину и высоту")
	fmt.Scanln(&width, &height)
	fmt.Printf("Площадь прямоугольника равна - %d", width*height)
}
