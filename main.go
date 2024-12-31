package main

import (
	"fmt"
)

func main() {
	head := NewLinkedList()
	head.InsertNodeAtFirst(4)
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNode(5)
	head.InsertNode(6)
	head.InsertNode(7)
	head.InsertNodeAtPosition(6, 10)
	head.InsertNodeAtMiddle(9)
	head.DeleteNodeByValue(5)
	head.DeleteNodeAtPosition(2)
	head.DeleteNode()
	head.ListAll()
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
