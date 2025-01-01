package main

import (
	"testing"
)

func TestInsertNodeInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)

	if doublyLinkedList.ToString() != "1" || doublyLinkedList.size != 1 {
		t.Fatalf("unable to add a node to LinkedList")
	}
}

func TestInsertAtFirstInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNodeAtFirst(3)

	if doublyLinkedList.ToString() != "3, 1" || doublyLinkedList.size != 2 {
		t.Fatalf("unable to add node to the first position of LinkedList")
	}
}

func TestInsertAtMiddleInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)

	if doublyLinkedList.ToString() != "1, 9, 3" || doublyLinkedList.size != 3 {
		t.Fatalf("unable to add a node to middle of LinkedList")
	}
}

func TestInsertAtPositionInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)

	if doublyLinkedList.ToString() != "1, 10, 9, 3" || doublyLinkedList.size != 4 {
		t.Fatalf("unable to add a node at position of LinkedList")
	}
}

func TestDeleteNodeInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	doublyLinkedList.DeleteNode()

	if doublyLinkedList.ToString() != "1, 10, 9" || doublyLinkedList.size != 3 {
		t.Fatalf("unable to delete a node from LinkedList")
	}
}

func TestDeleteAtPositionInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	doublyLinkedList.DeleteNodeAtPosition(1)

	if doublyLinkedList.ToString() != "1, 9, 3" || doublyLinkedList.size != 3 {
		t.Fatalf("unable to delete the node from the selected position in the LinkedList")
	}
}

func TestDeleteFirstInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	doublyLinkedList.DeleteFirstNode()

	if doublyLinkedList.ToString() != "10, 9, 3" || doublyLinkedList.size != 3 {
		t.Fatalf("unable to delete the first node from LinkedList")
	}
}

func TestDeleteByValueInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	doublyLinkedList.DeleteNodeByValue(9)

	if doublyLinkedList.ToString() != "1, 10, 3" || doublyLinkedList.size != 3 {
		t.Fatalf("unable to delete a node by value from LinkedList")
	}
}

func TestSearchInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)

	if !doublyLinkedList.Search(9) {
		t.Fatalf("unable to find the value from a node in the LinkedList")
	}
}

func TestSearchByPositionInDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	value, err := doublyLinkedList.SearchByPosition(2)

	if value != 9 || err != nil {
		t.Fatalf("unable to find the value from a node in the LinkedList")
	}
}

func TestSumOfDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)

	if doublyLinkedList.Sum() != 23 {
		t.Fatalf("unable to sum the values of the LinkedList")
	}
}

func TestDoublyLinkedListToArray(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	arr := doublyLinkedList.ToArray()

	if len(arr) != 4 {
		t.Fatalf("unable to convert the Doubly Linked List to array")
	}
}

func TestClearDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.InsertNodeAtPosition(1, 10)
	doublyLinkedList.Clear()

	if doublyLinkedList.size != 0 || doublyLinkedList.head != nil || doublyLinkedList.tail != nil {
		t.Fatalf("unable to clear the Doubly Linked List")
	}
}

func TestDeleteAtMiddleOnDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.DeleteNodeAtMiddle()

	if doublyLinkedList.ToString() != "1, 3" || doublyLinkedList.size != 2 {
		t.Fatalf("unable to delete node from the Doubly Linked List")
	}
}

func TestReverseDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.Reverse()

	if doublyLinkedList.ToString() != "3, 9, 1" || doublyLinkedList.size != 3 {
		t.Fatalf("unable to reverse doubly Linked List")
	}
}

func TestCheckHeadAndTailFromReversedDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()
	doublyLinkedList.InsertNode(1)
	doublyLinkedList.InsertNode(3)
	doublyLinkedList.InsertNodeAtMiddle(9)
	doublyLinkedList.Reverse()

	if doublyLinkedList.head.data != 3 {
		t.Fatalf("Wrong Head in reversed Doubly Linked List")
	}

	if doublyLinkedList.tail.data != 1 {
		t.Fatalf("Wrong Tail in reversed Doubly Linked List")
	}

	if doublyLinkedList.head.next.data != 9 {
		t.Fatalf("Wrong head next in reversed Doubly Linked List")
	}

	if doublyLinkedList.head.prev != nil {
		t.Fatalf("Head prev is not nil in reversed Doubly Linked List")
	}

	if doublyLinkedList.tail.prev.data != 9 {
		t.Fatalf("Wrong Tail prev in reversed Doubly Linked List")
	}

	if doublyLinkedList.tail.next != nil {
		t.Fatalf("Tail next is not nil in reversed Doubly Linked List")
	}
}
