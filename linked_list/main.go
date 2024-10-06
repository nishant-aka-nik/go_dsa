//https://www.educative.io/courses/data-structures-coding-interviews-python/singly-linked-lists-sll
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
	ll := LinkedList{}

	//displaying linkedlist when it is empty
	ll.Display_All_LinkedList_Nodes()

	ll.Insert_Data_At_Head(1)
	ll.Insert_Data_At_Head(2)
	ll.Insert_Data_At_Head(3)
	ll.Insert_Data_At_Head(4)

	//displaying linkedlist when it is not empty
	ll.Display_All_LinkedList_Nodes()

	//delete at head
	ll.Delete_at_head()

	//displaying linkedlist after deleting from head
	ll.Display_All_LinkedList_Nodes()
}

func (ll *LinkedList) Insert_Data_At_Head(data int) LinkedList {
	temp_node := Node{
		Data:     data,
		NextNode: ll.Head,
	}

	// ye jo temp node hai
	// ye kya krra hai ki jo new value aai
	// to vo kaise hogi
	// ye aisi hogi ki
	// data ki jagah data hoga
	// Next node ki jagah hum head ko daal denge kyuki head aage chala gya to next node ko head bana denge

	ll.Head = &temp_node

	// head ho gya next node
	// temp node ko head bana do
	return *ll
}

func (ll *LinkedList) Display_All_LinkedList_Nodes() {
	if ll.Head == nil {
		fmt.Println("linked list is empty")
		return
	}

	currentNode := ll.Head

	// fmt.Println("Head data -", currentNode.Data)

	// currentNode = currentNode.NextNode

	// fmt.Println("Next node data -", currentNode.Data)

	// currentNode = currentNode.NextNode

	// fmt.Println("Next node data -", currentNode.Data)

	// // yahi print hota rahega jab tak current ka next node nil ni hota

	// // for loop version v01 this has a bug that it will not print the last node as I am checking NextNode != nil
	// for currentNode.NextNode != nil {
	// 	fmt.Println(currentNode.Data)
	// 	currentNode = currentNode.NextNode
	// }

	// // iss loop me bhi same bug hai kyuki m current node ko ni next node ko check krra hu
	// // to last element jo hai vo print ni hora usse pehle he break ho ja ra hai
	// for {
	// 	fmt.Println(currentNode.Data)
	// 	currentNode = currentNode.NextNode
	// 	if currentNode.NextNode == nil {
	// 		break
	// 	}
	// }

	fmt.Println("Start linkedlist printing")
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.NextNode
	}
	fmt.Println("End linkedlist printing")
}

// There are three basic delete operations for linked lists:
// Deletion at the head
// Deletion by value
// Deletion at the tail
func (ll *LinkedList) Delete_at_head() {
	if ll.Head == nil {
		fmt.Println("linked list is empty")
		return
	}

	ll.Head = ll.Head.NextNode
}

func (ll *LinkedList) Delete_node_by_value(value int) {
	if ll.Head == nil {
		fmt.Println("linkedlist is empty")
	}

	
}
