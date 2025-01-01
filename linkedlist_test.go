package main

import (
	"testing"
)

func TestInsertNodeInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)

	if linkedList.ToString() != "1" || linkedList.size != 1 {
		t.Fatalf("unable to add a node to LinkedList")
	}
}

func TestInsertAtFirstInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNodeAtFirst(3)

	if linkedList.ToString() != "3, 1" || linkedList.size != 2 {
		t.Fatalf("unable to add node to the first position of LinkedList")
	}
}

func TestInsertAtMiddleInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)

	if linkedList.ToString() != "1, 9, 3" || linkedList.size != 3 {
		t.Fatalf("unable to add a node to middle of LinkedList")
	}
}

func TestInsertAtPositionInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)

	if linkedList.ToString() != "1, 10, 9, 3" || linkedList.size != 4 {
		t.Fatalf("unable to add a node at position of LinkedList")
	}
}

func TestDeleteNodeInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	linkedList.DeleteNode()

	if linkedList.ToString() != "1, 10, 9" || linkedList.size != 3 {
		t.Fatalf("unable to delete a node from LinkedList")
	}
}

func TestDeleteAtPositionInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	linkedList.DeleteNodeAtPosition(1)

	if linkedList.ToString() != "1, 9, 3" || linkedList.size != 3 {
		t.Fatalf("unable to delete the node from the selected position in the LinkedList")
	}
}

func TestDeleteFirstInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	linkedList.DeleteFirstNode()

	if linkedList.ToString() != "10, 9, 3" || linkedList.size != 3 {
		t.Fatalf("unable to delete the first node from LinkedList")
	}
}

func TestDeleteByValueInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	linkedList.DeleteNodeByValue(9)

	if linkedList.ToString() != "1, 10, 3" || linkedList.size != 3 {
		t.Fatalf("unable to delete a node by value from LinkedList")
	}
}

func TestSearchInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)

	if !linkedList.Search(9) {
		t.Fatalf("unable to find the value from a node in the LinkedList")
	}
}

func TestSearchByPositionInLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	value, err := linkedList.SearchByPosition(2)

	if value != 9 || err != nil {
		t.Fatalf("unable to find the value from a node in the LinkedList")
	}
}

func TestSumOfLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)

	if linkedList.Sum() != 23 {
		t.Fatalf("unable to sum the values of the LinkedList")
	}
}

func TestLinkedListToArray(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	arr := linkedList.ToArray()

	if len(arr) != 4 {
		t.Fatalf("unable to convert the LinkedList to array")
	}
}

func TestClearLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.InsertNodeAtPosition(1, 10)
	linkedList.Clear()

	if linkedList.size != 0 || linkedList.head != nil {
		t.Fatalf("unable to clear the LinkedList")
	}
}

func TestDeleteAtMiddleOfLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.DeleteNodeAtMiddle()

	if linkedList.ToString() != "1, 3" || linkedList.size != 2 {
		t.Fatalf("unable to delete node from the LinkedList")
	}
}

func TestReverseLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertNode(1)
	linkedList.InsertNode(3)
	linkedList.InsertNodeAtMiddle(9)
	linkedList.Reverse()

	if linkedList.ToString() != "3, 9, 1" || linkedList.size != 3 {
		t.Fatalf("unable to reverse LinkedList")
	}
}
