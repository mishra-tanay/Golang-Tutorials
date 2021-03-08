package main

import "fmt"

type Node struct {
	a int
	b string
}

func printNode(node Node) {
	fmt.Println("Index is ", node.a, "Name is ", node.b)
}

func setNode(node *Node) {
	node.a = 1
	node.b = "Tanay Mishra"
}

func main() {
	temp := Node{a: 1, b: "Mishra"}
	printNode(temp)
	printNode(Node{1, "Tanay"})
}
