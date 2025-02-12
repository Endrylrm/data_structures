package main

import "testing"

func TestPushNodeInStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)

	if stack.IsEmpty() {
		t.Fatalf("unable to push a node in the Stack")
	}
}

func TestPopNodeInStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Pop()

	if stack.ToString() != "3, 2, 1" || stack.size != 3 {
		t.Fatalf("unable to pop a node from the Stack")
	}
}

func TestPeekNodeInStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	node := stack.Peek()

	if node.data != 4 || node == nil {
		t.Fatalf("unable to peek the top node from the Stack")
	}
}

func TestClearTheStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Clear()

	if !stack.IsEmpty() {
		t.Fatalf("unable to clear the Stack")
	}
}

func TestReverseTheStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Reverse()

	if stack.ToString() != "1, 2, 3, 4" {
		t.Fatalf("unable to reverse the Stack")
	}
}

func TestSearchStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if !stack.Search(1) {
		t.Fatalf("unable to search the Stack")
	}
}
