package main

import (
	"fmt"
)

type Node[V comparable] struct {
	value       V
	left, right *Node[V]
}

func NewNode(value int) *Node[int] {
    return &Node[int]{value, nil, nil}
}

func (node *Node[V]) dfs(needle V) bool {
	if node == nil {
		return false
	}

	if node.value == needle {
		return true
	}

	return node.left.dfs(needle) || node.right.dfs(needle)
}

func main() {
	root := NewNode(3)
	root.left = NewNode(5)
	root.right = NewNode(8)
	root.right.left = NewNode(7)
	root.right.right = NewNode(0)
	root.right.left.right = NewNode(4)
    
    fmt.Println(root.right.left.value)
    fmt.Println("Finding 8... ")
    fmt.Println(root.dfs(8))
    fmt.Println("Finding 989... ")
    fmt.Println(root.dfs(989))
    fmt.Println("Finding 4... ")
    fmt.Println(root.dfs(4))
}
