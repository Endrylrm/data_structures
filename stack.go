package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (st *Stack) Size() int {
	return st.size
}

func (st *Stack) Clear() {
	for st.top != nil {
		temp := st.top.next
		st.top = nil
		st.top = temp
		st.size--
	}
}

func (st *Stack) IsEmpty() bool {
	return st.size == 0 && st.top == nil
}

func (st *Stack) Peek() *Node {
	return st.top
}

func (st *Stack) Pop() {
	if st.top == nil {
		fmt.Println("The stack is empty!")
		return
	}

	st.top = st.top.next
	st.size--
}

func (st *Stack) Push(value int) {
	newNode := &Node{value, nil}

	if st.top == nil {
		st.top = newNode
	} else {
		newNode.next = st.top
		st.top = newNode
	}
	st.size++
}

func (st *Stack) Reverse() {
	var prevNode, nextNode *Node
	curNode := st.top

	for curNode != nil {
		nextNode = curNode.next
		curNode.next = prevNode
		prevNode = curNode
		curNode = nextNode
	}

	st.top = prevNode
}

func (st *Stack) Search(value int) bool {
	curNode := st.top
	distance := 0

	for curNode != nil {
		if curNode.data == value {
			fmt.Println("Distance to value: ", distance)
			return true
		}
		distance++
		curNode = curNode.next
	}

	return false
}

func (st *Stack) SearchByPosition(index int) (int, error) {
	curNode := st.top
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

func (st *Stack) ToString() string {
	curNode := st.top
	var dataList []string

	defer clear(dataList)

	for curNode != nil {
		dataList = append(dataList, strconv.Itoa(curNode.data))
		curNode = curNode.next
	}

	allDataString := strings.Join(dataList, ", ")

	return allDataString
}

func (st *Stack) ToArray() []int {
	curNode := st.top
	var dataList []int

	for curNode != nil {
		dataList = append(dataList, curNode.data)
		curNode = curNode.next
	}

	return dataList
}

func (st *Stack) Traverse() {
	curNode := st.top

	for curNode != nil {
		fmt.Print(curNode.data, " => ")
		curNode = curNode.next
	}

	fmt.Println("nil")
}
