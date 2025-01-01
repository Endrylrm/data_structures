package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DoublyNode struct {
	data int
	next *DoublyNode
	prev *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
	size int
}

func (dll *DoublyLinkedList) CheckValidPosition(index int) bool {
	if index > dll.size-1 {
		fmt.Println("Out of bounds, please put a valid value")
		return false
	}
	return true
}

func (dll *DoublyLinkedList) Clear() {
	for dll.head != nil {
		temp := dll.head.next
		dll.head = nil
		dll.head = temp
		dll.size--
	}
	if dll.size == 0 {
		dll.tail = nil
	}
}

func (dll *DoublyLinkedList) DeleteFirstNode() {
	dll.head = dll.head.next
	dll.head.prev = nil

	if dll.head == nil {
		dll.tail = nil
	}

	dll.size--
}

func (dll *DoublyLinkedList) DeleteNode() {
	if dll.size == 1 && dll.head != nil {
		dll.DeleteFirstNode()
		return
	}

	dll.tail = dll.tail.prev
	dll.tail.next = nil
	dll.size--
}

func (dll *DoublyLinkedList) DeleteNodeAtMiddle() {
	dll.DeleteNodeAtPosition(dll.size / 2)
}

func (dll *DoublyLinkedList) DeleteNodeAtPosition(index int) {
	if !dll.CheckValidPosition(index) {
		return
	}

	if index == 0 {
		dll.DeleteFirstNode()
		return
	}

	curNode := dll.head
	prevNode := &DoublyNode{0, nil, nil}

	count := 0
	for curNode != nil && count != index {
		prevNode = curNode
		curNode = curNode.next
		count++
	}

	if curNode != nil {
		prevNode.next = curNode.next
		curNode.next.prev = prevNode
		dll.size--
	}
}

func (dll *DoublyLinkedList) DeleteNodeByValue(value int) {
	if dll.head != nil && dll.head.data == value {
		dll.DeleteFirstNode()
		return
	}

	curNode := dll.head
	prevNode := &DoublyNode{0, nil, nil}

	for curNode != nil && curNode.data != value {
		prevNode = curNode
		curNode = curNode.next
	}

	if curNode != nil {
		prevNode.next = curNode.next
		curNode.next.prev = prevNode
		dll.size--
	}
}

func (dll *DoublyLinkedList) InsertNode(value int) {
	newNode := &DoublyNode{value, nil, nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) InsertNodeAtFirst(value int) {
	if dll.head == nil {
		dll.InsertNode(value)
		return
	}

	curHead := dll.head
	newHead := &DoublyNode{value, nil, nil}
	newHead.next = curHead
	curHead.prev = newHead
	dll.head = newHead
	dll.size++
}

func (dll *DoublyLinkedList) InsertNodeAtMiddle(value int) {
	dll.InsertNodeAtPosition(dll.size/2, value)
}

func (dll *DoublyLinkedList) InsertNodeAtPosition(index int, value int) {
	if index == dll.size {
		dll.InsertNode(value)
		return
	}

	if !dll.CheckValidPosition(index) {
		return
	}

	curNode := dll.head
	prevNode := &DoublyNode{0, nil, nil}

	count := 0
	for curNode != nil && count != index {
		prevNode = curNode
		curNode = curNode.next
		count++
	}

	if curNode != nil {
		newNode := &DoublyNode{value, curNode, prevNode}
		prevNode.next = newNode
		curNode.prev = newNode
		dll.size++
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

func (dll *DoublyLinkedList) Reverse() {
	curNode := dll.head

	for curNode != nil {
		curNode.next, curNode.prev = curNode.prev, curNode.next
		curNode = curNode.prev
	}

	dll.head, dll.tail = dll.tail, dll.head
}

func (dll *DoublyLinkedList) Search(value int) bool {
	curNode := dll.head

	for curNode != nil {
		if curNode.data == value {
			return true
		}
		curNode = curNode.next
	}

	return false
}

func (dll *DoublyLinkedList) SearchByPosition(index int) (int, error) {
	curNode := dll.head
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

func (dll *DoublyLinkedList) Sum() int {
	sum := 0
	curNode := dll.head

	for curNode != nil {
		sum += curNode.data
		curNode = curNode.next
	}

	return sum
}

func (dll *DoublyLinkedList) ToString() string {
	curNode := dll.head
	var dataList []string

	defer clear(dataList)

	for curNode != nil {
		dataList = append(dataList, strconv.Itoa(curNode.data))
		curNode = curNode.next
	}

	allDataString := strings.Join(dataList, ", ")

	return allDataString
}

func (dll *DoublyLinkedList) ToStringReverse() string {
	curNode := dll.tail
	var dataList []string

	defer clear(dataList)

	for curNode != nil {
		dataList = append(dataList, strconv.Itoa(curNode.data))
		curNode = curNode.prev
	}

	allDataString := strings.Join(dataList, ", ")

	return allDataString
}

func (dll *DoublyLinkedList) ToArray() []int {
	curNode := dll.head
	var dataList []int

	for curNode != nil {
		dataList = append(dataList, curNode.data)
		curNode = curNode.next
	}

	return dataList
}

func (dll *DoublyLinkedList) ToArrayReverse() []int {
	curNode := dll.tail
	var dataList []int

	for curNode != nil {
		dataList = append(dataList, curNode.data)
		curNode = curNode.prev
	}

	return dataList
}
