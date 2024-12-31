package main

import (
	"testing"
)

func TestInsertNewNode(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)

	if head.ToString() != "1" || head.size != 1 {
		t.Fatalf("unable to add a node to LinkedList")
	}
}

func TestInsertAtFirst(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNodeAtFirst(3)

	if head.ToString() != "3, 1" || head.size != 2 {
		t.Fatalf("unable to add node to the first position of LinkedList")
	}
}

func TestInsertAtMiddle(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)

	if head.ToString() != "1, 9, 3" || head.size != 3 {
		t.Fatalf("unable to add a node to middle of LinkedList")
	}
}

func TestInsertAtPosition(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)

	if head.ToString() != "1, 10, 9, 3" || head.size != 4 {
		t.Fatalf("unable to add a node at position of LinkedList")
	}
}

func TestDeleteNode(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)
	head.DeleteNode()

	if head.ToString() != "1, 10, 9" || head.size != 3 {
		t.Fatalf("unable to delete a node from LinkedList")
	}
}

func TestDeleteAtPosition(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)
	head.DeleteNodeAtPosition(1)

	if head.ToString() != "1, 9, 3" || head.size != 3 {
		t.Fatalf("unable to delete the node from the selected position in the LinkedList")
	}
}

func TestDeleteFirst(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)
	head.DeleteFirstNode()

	if head.ToString() != "10, 9, 3" || head.size != 3 {
		t.Fatalf("unable to delete the first node from LinkedList")
	}
}

func TestDeleteByValue(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)
	head.DeleteNodeByValue(9)

	if head.ToString() != "1, 10, 3" || head.size != 3 {
		t.Fatalf("unable to delete a node by value from LinkedList")
	}
}

func TestSearch(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)

	if !head.Search(9) {
		t.Fatalf("unable to find the value from a node in the LinkedList")
	}
}

func TestSearchByPosition(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)
	value, err := head.SearchByPosition(2)

	if value != 9 || err != nil {
		t.Fatalf("unable to find the value from a node in the LinkedList")
	}
}

func TestSum(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)

	if head.Sum() != 23 {
		t.Fatalf("unable to sum the values of the LinkedList")
	}
}

func TestLinkedListToArray(t *testing.T) {
	head := NewLinkedList()
	head.InsertNode(1)
	head.InsertNode(3)
	head.InsertNodeAtMiddle(9)
	head.InsertNodeAtPosition(1, 10)
	arr := head.ToArray()

	if len(arr) != 4 {
		t.Fatalf("unable to convert the LinkedList to array")
	}
}
