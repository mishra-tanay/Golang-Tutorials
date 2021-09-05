package main

import (
	"fmt"
)

var list []string

func printValue(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, " is printing ", i)
		list = append(list, from)
	}
}
func main() {

	messages := make(chan string)
	go func() { messages <- "Called" }()
	go printValue("Mishra")

	printValue("Tanu")
	msgs := <-messages
	fmt.Println(msgs)
}
