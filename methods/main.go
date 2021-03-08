package main

import "fmt"

type Student struct {
	age  int
	name string
}

func (s *Student) printInfo() {
	fmt.Println("Name of student is ", s.name, "and age is ", s.age)
}
func main() {
	a := Student{19, "Tanay Mishra"}
	a.printInfo()
}
