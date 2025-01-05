package main

import "testing"

func TestEnqueueNodeInQueue(t *testing.T) {
	queue := NewQueue(1)
	queue.Enqueue(1)

	if queue.IsEmpty() {
		t.Fatalf("unable to push a node in the Stack")
	}
}

func TestDequeueNodeInQueue(t *testing.T) {
	queue := NewQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()

	if queue.IsFull() {
		t.Fatalf("unable to push a node in the Stack")
	}
}

func TestQueueCapacity(t *testing.T) {
	queue := NewQueue(1)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if !queue.IsFull() {
		t.Fatalf("unable to push a node in the Stack")
	}
}
