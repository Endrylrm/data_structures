package main

import "testing"

func TestEnqueueNodeInQueue(t *testing.T) {
	queue := NewQueue(1)
	queue.Enqueue(1)

	if queue.IsEmpty() {
		t.Fatalf("unable to enqueue a node in the Queue")
	}
}

func TestDequeueNodeInQueue(t *testing.T) {
	queue := NewQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()

	if queue.IsFull() {
		t.Fatalf("unable to dequeue a node in the Queue")
	}
}

func TestQueueCapacity(t *testing.T) {
	queue := NewQueue(1)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if queue.Size() > 1 {
		t.Fatalf("Size of Queue is bigger than it's capacity")
	}
}

func TestClearQueue(t *testing.T) {
	queue := NewQueue(1)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Clear()

	if !queue.IsEmpty() {
		t.Fatalf("unable to clear the Queue")
	}
}
