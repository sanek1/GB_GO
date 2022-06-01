package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&a)
	fmt.Printf("Количество сотен - %d\n", a/100)
	fmt.Printf("Количество десятков - %d\n", a%100/10)
	fmt.Printf("Количество  единиц- %d\n", a%100%10)
}
