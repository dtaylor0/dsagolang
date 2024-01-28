package main

import (
	"fmt"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func newNode(value int) *Node {
	return &Node{value, nil, nil}
}

type DoublyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func newDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{0, nil, nil}
}

func prepend(lst *DoublyLinkedList, item int) {
	node := newNode(item)
	lst.length++
	if lst.head == nil {
		lst.head = node
		lst.tail = node
		return
	}
	node.next = lst.head
	lst.head.prev = node
	lst.head = node
}

func insertAt(lst *DoublyLinkedList, item int, idx int) {
	if idx > lst.length {
		panic("Bad index")
	} else if idx == lst.length {
		appendll(lst, item)
		return
	} else if idx == 0 {
		prepend(lst, item)
		return
	}
	lst.length++
	curr := lst.head
	for idx > 0 {
		idx--
		curr = curr.next
	}
	node := newNode(item)

	node.prev = curr.prev
	node.next = curr

	curr.prev.next = node
	if curr.prev != nil {
		curr.prev = node
	}
}

func appendll(lst *DoublyLinkedList, item int) {
	lst.length++
	node := newNode(item)
	if lst.tail == nil {
		lst.head = node
		lst.tail = node
		return
	}
	node.prev = lst.tail
	lst.tail.next = node
	lst.tail = node
}

func remove(lst *DoublyLinkedList, item int) (int, bool) {
	curr := lst.head
	found := false
	for i := 0; i < lst.length; i++ {
		if curr.value == item {
			found = true
			break
		}
		curr = curr.next
	}

	if !found {
		return -1, false
	}
	removeNode(lst, curr)
	return curr.value, true
}

func removeAt(lst *DoublyLinkedList, idx int) (int, bool) {
	node := getAt(lst, idx)
	if node == nil {
		return -1, false
	}
	removeNode(lst, node)
	return node.value, true
}

func getAt(lst *DoublyLinkedList, idx int) *Node {
	if idx >= lst.length {
		panic("Bad index")
	}
	curr := lst.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr
}

func removeNode(lst *DoublyLinkedList, node *Node) *Node {
	lst.length--
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lst.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lst.tail = node.prev
	}
	return node
}

func printList(lst *DoublyLinkedList) {
	curr := lst.head
	fmt.Printf("Length: %d ", lst.length)
	fmt.Print("[")
	for curr != nil {
		fmt.Printf("%d, ", curr.value)
		curr = curr.next
	}
	fmt.Println("]")
}

func main() {
	test := newDoublyLinkedList()
	printList(test)
	fmt.Println("Appending 1")
	appendll(test, 1)
	printList(test)
	fmt.Println("Appending 2")
	appendll(test, 2)
	printList(test)
	fmt.Println("Appending 3")
	appendll(test, 3)
	printList(test)
	fmt.Println("Prepending 0")
	prepend(test, 0)
	printList(test)
	fmt.Println("Inserting 4 at index 4")
	insertAt(test, 4, 4)
	printList(test)
	fmt.Println("Inserting 5 at index 0")
	insertAt(test, 5, 0)
	printList(test)
	fmt.Println("Inserting 6 at index 3")
	insertAt(test, 6, 3)
	printList(test)
	fmt.Println("Removing at index 0")
	removeAt(test, 0)
	printList(test)
	fmt.Println("Removing at index 5")
	removeAt(test, 5)
	printList(test)
	fmt.Println("Removing at index 3")
	removeAt(test, 3)
	printList(test)
	fmt.Println("Removing 3")
	remove(test, 3)
	printList(test)
	fmt.Println("Removing 0")
	remove(test, 0)
	printList(test)
	fmt.Println("Removing 6")
	remove(test, 6)
	printList(test)
	fmt.Println("Removing 1")
	remove(test, 1)
	printList(test)
	fmt.Println("Removing 2")
	remove(test, 2)
	printList(test)
}
