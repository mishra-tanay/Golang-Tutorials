package main

import "fmt"

func passByValue(num int) {
	num++
	fmt.Println("NUM in func", num)
}

func passByReference(num *int) {
	*num++
	fmt.Println("Num in func", *num)
}
func main() {
	var i int = 2
	passByValue(i)
	fmt.Println("Num in main ", i)
	passByReference(&i)
	fmt.Println("Num in main", i)
}
