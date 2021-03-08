package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Tanay"] = 1
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "Tanay")
	fmt.Println(m)
	_, ok := m["Tanay"]
	if ok {
		fmt.Println("Key present")
	} else {
		fmt.Println("Key not present")
	}

	mp := map[string]int{"Tanay": 2, "Mishra": 1}
	fmt.Println(mp)
	var mp1 map[string]int
	fmt.Println(mp1)
}
