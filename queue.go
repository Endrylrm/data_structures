package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Queue struct {
	front    *Node
	rear     *Node
	capacity int
	size     int
}

func NewQueue(capacity int) *Queue {
	return &Queue{nil, nil, capacity, 0}
}

func (queue *Queue) Dequeue() {
	if queue.IsEmpty() {
		fmt.Println("No nodes to Dequeue!")
		return
	}

	queue.front = queue.front.next

	if queue.front == nil {
		queue.rear = nil
	}

	queue.size--
}

func (queue *Queue) Enqueue(value int) {
	if queue.IsFull() {
		fmt.Println("Queue is full!")
		return
	}

	newNode := &Node{value, nil}

	if queue.front == nil && queue.rear == nil {
		queue.front = newNode
		queue.rear = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}

	queue.size++
}

func (queue *Queue) Clear() {
	for queue.front != nil {
		temp := queue.front.next
		queue.front = nil
		queue.front = temp
		queue.size--
	}

	queue.rear = nil
}

func (queue *Queue) Front() *Node {
	return queue.front
}

func (queue *Queue) Rear() *Node {
	return queue.rear
}

func (queue *Queue) Capacity() int {
	return queue.capacity
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0 && queue.front == nil && queue.rear == nil
}

func (queue *Queue) IsFull() bool {
	return queue.size == queue.capacity
}

func (queue *Queue) Reverse() {
	var prevNode, nextNode *Node
	curNode := queue.front

	for curNode != nil {
		nextNode = curNode.next
		curNode.next = prevNode
		prevNode = curNode
		curNode = nextNode
	}

	queue.rear = queue.front
	queue.front = prevNode
}

func (queue *Queue) Search(value int) bool {
	curNode := queue.front

	for curNode != nil {
		if curNode.data == value {
			return true
		}
		curNode = curNode.next
	}

	return false
}

func (queue *Queue) SearchByPosition(index int) (int, error) {
	curNode := queue.front
	count := 0

	for curNode != nil {
		if count == index {
			return curNode.data, nil
		}
		curNode = curNode.next
		count++
	}

	return 0, errors.New("out of bounds error")
}

func (queue *Queue) ToString() string {
	curNode := queue.front
	var dataList []string

	defer clear(dataList)

	for curNode != nil {
		dataList = append(dataList, strconv.Itoa(curNode.data))
		curNode = curNode.next
	}

	allDataString := strings.Join(dataList, ", ")

	return allDataString
}

func (queue *Queue) ToArray() []int {
	curNode := queue.front
	var dataList []int

	for curNode != nil {
		dataList = append(dataList, curNode.data)
		curNode = curNode.next
	}

	return dataList
}

func (queue *Queue) Traverse() {
	curNode := queue.front

	for curNode != nil {
		fmt.Print(curNode.data, " => ")
		curNode = curNode.next
	}

	fmt.Println("nil")
}
