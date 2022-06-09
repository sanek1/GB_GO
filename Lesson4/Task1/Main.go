package main

import (
	"fmt"
)

func main() {

	slice := []int{314, 620, 559, 396, 549, -726, 34, -9, 629, 696, -532, 89, 204, -371, -161, 466, 21, 34, 335, 188}
	fmt.Println("Unsorted\n", slice)
	insertionsort(slice)
	fmt.Println("Sorted\n", slice, "")
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
