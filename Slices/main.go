package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "ABC"
	s[1] = "BCS"

	fmt.Println(s, len(s))
	s = append(s, "tanay", "Mishra")
	fmt.Println(s, len(s))

}
