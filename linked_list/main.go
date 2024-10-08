// https://www.educative.io/courses/data-structures-coding-interviews-python/singly-linked-lists-sll
package main

import (
	"errors"
	"fmt"
)

// type Node struct {
// 	Data     int
// 	NextNode *Node
// }

// type LinkedList struct {
// 	Head *Node
// }

// func main() {
// 	ll := LinkedList{}

// 	//displaying linkedlist when it is empty
// 	ll.Display_All_LinkedList_Nodes()

// 	ll.Insert_Data_At_Head(1)
// 	ll.Insert_Data_At_Head(2)
// 	ll.Insert_Data_At_Head(3)
// 	ll.Insert_Data_At_Head(4)

// 	//displaying linkedlist when it is not empty
// 	ll.Display_All_LinkedList_Nodes()

// 	//delete at head
// 	ll.Delete_at_head()

// 	//displaying linkedlist after deleting from head
// 	ll.Display_All_LinkedList_Nodes()
// }

// func (ll *LinkedList) Insert_Data_At_Head(data int) LinkedList {
// 	temp_node := Node{
// 		Data:     data,
// 		NextNode: ll.Head,
// 	}

// 	// ye jo temp node hai
// 	// ye kya krra hai ki jo new value aai
// 	// to vo kaise hogi
// 	// ye aisi hogi ki
// 	// data ki jagah data hoga
// 	// Next node ki jagah hum head ko daal denge kyuki head aage chala gya to next node ko head bana denge

// 	ll.Head = &temp_node

// 	// head ho gya next node
// 	// temp node ko head bana do
// 	return *ll
// }

// func (ll *LinkedList) Display_All_LinkedList_Nodes() {
// 	if ll.Head == nil {
// 		fmt.Println("linked list is empty")
// 		return
// 	}

// 	currentNode := ll.Head

// 	// fmt.Println("Head data -", currentNode.Data)

// 	// currentNode = currentNode.NextNode

// 	// fmt.Println("Next node data -", currentNode.Data)

// 	// currentNode = currentNode.NextNode

// 	// fmt.Println("Next node data -", currentNode.Data)

// 	// // yahi print hota rahega jab tak current ka next node nil ni hota

// 	// // for loop version v01 this has a bug that it will not print the last node as I am checking NextNode != nil
// 	// for currentNode.NextNode != nil {
// 	// 	fmt.Println(currentNode.Data)
// 	// 	currentNode = currentNode.NextNode
// 	// }

// 	// // iss loop me bhi same bug hai kyuki m current node ko ni next node ko check krra hu
// 	// // to last element jo hai vo print ni hora usse pehle he break ho ja ra hai
// 	// for {
// 	// 	fmt.Println(currentNode.Data)
// 	// 	currentNode = currentNode.NextNode
// 	// 	if currentNode.NextNode == nil {
// 	// 		break
// 	// 	}
// 	// }

// 	fmt.Println("Start linkedlist printing")
// 	for currentNode != nil {
// 		fmt.Println(currentNode.Data)
// 		currentNode = currentNode.NextNode
// 	}
// 	fmt.Println("End linkedlist printing")
// }

// // There are three basic delete operations for linked lists:
// // Deletion at the head
// // Deletion by value
// // Deletion at the tail
// func (ll *LinkedList) Delete_at_head() {
// 	if ll.Head == nil {
// 		fmt.Println("linked list is empty")
// 		return
// 	}

// 	ll.Head = ll.Head.NextNode
// }

// func (ll *LinkedList) Delete_node_by_value(value int) {
// 	if ll.Head == nil {
// 		fmt.Println("linkedlist is empty")
// 	}

// }

// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents the singly linked list
type LinkedList struct {
	Head *Node
}

// InsertAtBeginning adds a new node at the start of the list
func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
}

// InsertAtEnd adds a new node at the end of the list
func (ll *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{Data: data, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// InsertAtPosition inserts a new node at the specified position (0-based index)
func (ll *LinkedList) InsertAtPosition(data int, position int) error {
	if position < 0 {
		return errors.New("position cannot be negative")
	}

	newNode := &Node{Data: data, Next: nil}

	if position == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return nil
	}

	current := ll.Head
	for i := 0; i < position-1; i++ {
		if current == nil {
			return errors.New("position out of bounds")
		}
		current = current.Next
	}

	if current == nil {
		return errors.New("position out of bounds")
	}

	newNode.Next = current.Next
	current.Next = newNode
	return nil
}

// DeleteAtBeginning removes the first node of the list
func (ll *LinkedList) DeleteAtBeginning() error {
	if ll.Head == nil {
		return errors.New("list is empty")
	}
	ll.Head = ll.Head.Next
	return nil
}

// DeleteAtEnd removes the last node of the list
func (ll *LinkedList) DeleteAtEnd() error {
	if ll.Head == nil {
		return errors.New("list is empty")
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		return nil
	}

	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
	return nil
}

// DeleteAtPosition removes a node from the specified position (0-based index)
func (ll *LinkedList) DeleteAtPosition(position int) error {
	if position < 0 {
		return errors.New("position cannot be negative")
	}

	if ll.Head == nil {
		return errors.New("list is empty")
	}

	if position == 0 {
		ll.Head = ll.Head.Next
		return nil
	}

	current := ll.Head
	for i := 0; i < position-1; i++ {
		if current.Next == nil {
			return errors.New("position out of bounds")
		}
		current = current.Next
	}

	if current.Next == nil {
		return errors.New("position out of bounds")
	}

	current.Next = current.Next.Next
	return nil
}

// Search finds the first node with the specified data and returns its position
func (ll *LinkedList) Search(data int) (int, bool) {
	current := ll.Head
	position := 0
	for current != nil {
		if current.Data == data {
			return position, true
		}
		current = current.Next
		position++
	}
	return -1, false
}

// Update modifies the data of the node at the specified position
func (ll *LinkedList) Update(position int, newData int) error {
	if position < 0 {
		return errors.New("position cannot be negative")
	}

	current := ll.Head
	for i := 0; i < position; i++ {
		if current == nil {
			return errors.New("position out of bounds")
		}
		current = current.Next
	}

	if current == nil {
		return errors.New("position out of bounds")
	}

	current.Data = newData
	return nil
}

// Traverse prints all the nodes in the list
func (ll *LinkedList) Traverse() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

// Reverse reverses the linked list in place
func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.Head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	ll.Head = prev
}

// DetectLoop checks if there's a cycle in the linked list using Floyd’s Cycle-Finding Algorithm
func (ll *LinkedList) DetectLoop() bool {
	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// Example usage
func main() {
	ll := LinkedList{}

	// Insertion
	ll.InsertAtEnd(10)
	ll.InsertAtBeginning(5)
	ll.InsertAtEnd(15)
	ll.InsertAtPosition(7, 1) // List: 5 -> 7 -> 10 -> 15
	fmt.Print("List after insertions: ")
	ll.Traverse()

	// Deletion
	ll.DeleteAtBeginning() // Removes 5
	fmt.Print("List after deleting at beginning: ")
	ll.Traverse()

	ll.DeleteAtEnd() // Removes 15
	fmt.Print("List after deleting at end: ")
	ll.Traverse()

	ll.DeleteAtPosition(0) // Removes 7
	fmt.Print("List after deleting at position 0: ")
	ll.Traverse()

	// Search
	pos, found := ll.Search(10)
	if found {
		fmt.Printf("Element 10 found at position %d\n", pos)
	} else {
		fmt.Println("Element 10 not found")
	}

	// Update
	err := ll.Update(0, 20) // Update first element to 20
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("List after update: ")
		ll.Traverse()
	}

	// Reverse
	ll.InsertAtEnd(30)
	ll.InsertAtEnd(40)
	fmt.Print("List before reversing: ")
	ll.Traverse()
	ll.Reverse()
	fmt.Print("List after reversing: ")
	ll.Traverse()

	// Detect Loop
	fmt.Printf("Does the list have a loop? %v\n", ll.DetectLoop())

	// Creating a loop for demonstration (not recommended in practice)
	// Uncomment the following lines to create a loop and detect it
	/*
		ll.Head.Next.Next.Next = ll.Head // Creates a loop
		fmt.Printf("Does the list have a loop? %v\n", ll.DetectLoop())
	*/
}
