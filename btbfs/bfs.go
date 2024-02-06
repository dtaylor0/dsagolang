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

func (haystack *Node[V]) bfs(needle V) bool {
	q := [](*Node[V]){haystack}
	for len(q) > 0 {
        curr := q[0]
        q = q[1:]
		if curr.value == needle {
			return true
		}
        if curr.left != nil {
            q = append(q, curr.left)
        }
        if curr.right != nil {
            q = append(q, curr.right)
        }
	}
	return false
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
    fmt.Println(root.bfs(8))
    fmt.Println("Finding 989... ")
    fmt.Println(root.bfs(989))
    fmt.Println("Finding 4... ")
    fmt.Println(root.bfs(4))
}
