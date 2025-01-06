package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

func (ll *LinkedList) Head() *Node {
	return ll.head
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) CheckValidPosition(index int) bool {
	if index > ll.size-1 {
		fmt.Println("Out of bounds, please put a valid value")
		return false
	}
	return true
}

func (ll *LinkedList) Clear() {
	for ll.head != nil {
		temp := ll.head.next
		ll.head = nil
		ll.head = temp
		ll.size--
	}
}

func (ll *LinkedList) DeleteFirstNode() {
	ll.head = ll.head.next
	ll.size--
}

func (ll *LinkedList) DeleteNode() {
	if ll.size == 1 && ll.head != nil {
		ll.DeleteFirstNode()
		return
	}

	count := 0

	curHead := ll.head
	for curHead.next != nil && count != ll.size-2 {
		curHead = curHead.next
		count++
	}

	curHead.next = nil
	ll.size--
}

func (ll *LinkedList) DeleteNodeAtMiddle() {
	ll.DeleteNodeAtPosition(ll.size / 2)
}

func (ll *LinkedList) DeleteNodeAtPosition(index int) {
	if !ll.CheckValidPosition(index) {
		return
	}

	if index == 0 {
		ll.DeleteFirstNode()
		return
	}

	curNode := ll.head
	prevNode := &Node{0, nil}

	count := 0
	for curNode != nil && count != index {
		prevNode = curNode
		curNode = curNode.next
		count++
	}

	if curNode != nil {
		prevNode.next = curNode.next
		ll.size--
	}
}

func (ll *LinkedList) DeleteNodeByValue(value int) {
	if ll.head != nil && ll.head.data == value {
		ll.DeleteFirstNode()
		return
	}

	curNode := ll.head
	prevNode := &Node{0, nil}

	for curNode != nil && curNode.data != value {
		prevNode = curNode
		curNode = curNode.next
	}

	if curNode != nil {
		prevNode.next = curNode.next
		ll.size--
	}
}

func (ll *LinkedList) InsertNode(value int) {
	newNode := &Node{value, nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		curNode := ll.head

		for curNode.next != nil {
			curNode = curNode.next
		}

		curNode.next = newNode
	}
	ll.size++
}

func (ll *LinkedList) InsertNodeAtFirst(value int) {
	curHead := ll.head
	newHead := &Node{value, curHead}
	ll.head = newHead
	ll.size++
}

func (ll *LinkedList) InsertNodeAtMiddle(value int) {
	ll.InsertNodeAtPosition(ll.size/2, value)
}

func (ll *LinkedList) InsertNodeAtPosition(index int, value int) {
	if index == ll.size {
		ll.InsertNode(value)
		return
	}

	if !ll.CheckValidPosition(index) {
		return
	}

	curNode := ll.head
	prevNode := &Node{0, nil}

	count := 0
	for curNode != nil && count != index {
		prevNode = curNode
		curNode = curNode.next
		count++
	}

	if curNode != nil {
		newNode := &Node{value, nil}
		prevNode.next = newNode
		newNode.next = curNode
		ll.size++
	}
}

func (ll *LinkedList) Reverse() {
	var prevNode, nextNode *Node
	curNode := ll.head

	for curNode != nil {
		nextNode = curNode.next
		curNode.next = prevNode
		prevNode = curNode
		curNode = nextNode
	}

	ll.head = prevNode
}

func (ll *LinkedList) Search(value int) bool {
	curNode := ll.head

	for curNode != nil {
		if curNode.data == value {
			return true
		}
		curNode = curNode.next
	}

	return false
}

func (ll *LinkedList) SearchByPosition(index int) (int, error) {
	curNode := ll.head
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

func (ll *LinkedList) Sum() int {
	sum := 0
	curNode := ll.head

	for curNode != nil {
		sum += curNode.data
		curNode = curNode.next
	}

	return sum
}

func (ll *LinkedList) ToString() string {
	curNode := ll.head
	var dataList []string

	defer clear(dataList)

	for curNode != nil {
		dataList = append(dataList, strconv.Itoa(curNode.data))
		curNode = curNode.next
	}

	allDataString := strings.Join(dataList, ", ")

	return allDataString
}

func (ll *LinkedList) ToArray() []int {
	curNode := ll.head
	var dataList []int

	for curNode != nil {
		dataList = append(dataList, curNode.data)
		curNode = curNode.next
	}

	return dataList
}

func (ll *LinkedList) Traverse() {
	curNode := ll.head

	for curNode != nil {
		fmt.Print(curNode.data, " => ")
		curNode = curNode.next
	}

	fmt.Println("nil")
}
