package main

import "fmt"

type Node struct {
	Data     int
	NextNode *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	linkedList := CreateLinkedList()
	fmt.Println("linked list 1", linkedList)
}

func CreateLinkedList() LinkedList {
	linkedList := LinkedList{
		Head: nil,
	}
	return linkedList
}
