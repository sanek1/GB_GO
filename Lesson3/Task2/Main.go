package main

import (
	"fmt"
)

// алгоритм подсмотрел здесь
// https://www.techiedelight.com/ru/sieve-of-eratosthenes/
// не самый быстрый
func findPrimes(N int) (b []int) {
	for i := 0; i < N; i++ {
		b = append(b, 1)
	}
	for i := 2; i < N; i++ {
		if b[i] == 1 {
			for j := 2; i*j < N; j++ {
				b[i*j] = 0
			}
			continue
		}
	}
	return
}

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)
	primes := findPrimes(a)
	for i := 2; i < len(primes); i++ {
		if primes[i] == 1 {
			fmt.Println(i)
		}
	}
}
