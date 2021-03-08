package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("Array initialised ", a)
	for i := 0; i < len(a); i++ {
		a[i] = 2 * i
	}
	fmt.Println("Array changed to ", a)

	b := [5]int{1, 2, 3}
	fmt.Println("B array ", b)

	var twoD [2][2]int

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			fmt.Println(i, j, twoD[i][j])
		}
	}
}
