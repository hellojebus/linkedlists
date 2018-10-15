package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func main() {
	fmt.Println("vim-go")
}
