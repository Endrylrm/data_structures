package main

import "fmt"

func main() {
	head := NewDoublyLinkedList()
	head.InsertNodeAtFirst(4)
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNode(5)
	head.InsertNode(6)
	head.InsertNode(7)
	head.InsertNodeAtPosition(3, 10)
	head.InsertNodeAtMiddle(9)
	head.DeleteNodeByValue(5)
	head.DeleteNodeAtPosition(2)
	head.DeleteNode()
	head.DeleteNodeAtMiddle()
	head.Reverse()
	head.Traverse()
	fmt.Println(head.ToString())
	fmt.Println(head.ToStringReverse())
	arr := head.ToArray()
	value, err := head.SearchByPosition(1)
	fmt.Println("the value at index 1 is:", value)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current nodes:", head.size)
	fmt.Println("Current sum of LinkedList is:", head.Sum())
	fmt.Println(head.Search(9))
	fmt.Println(arr)
}
