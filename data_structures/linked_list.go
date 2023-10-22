package main

import "errors"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	dummy := &Node{}
	return &LinkedList{head: dummy}
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {
		currentNode := l.head

		// traverse to tail
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func (l *LinkedList) Insert(data int, index int) error {
	if index < 0 || index > l.Len() {
		return errors.New("Index out of bounds")
	}

	newNode := &Node{data: data}
	currentNode := l.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode
	return nil
}

func (l *LinkedList) Delete(index int) error {
	if index < 0 || index >= l.Len() {
		return errors.New("Index out of bounds")
	}

	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	currentNode.next = currentNode.next.next
	return nil
}

func (l *LinkedList) Get(index int) int {
	currentNode := l.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (l *LinkedList) Len() int {
	length := 0
	currentNode := l.head

	for currentNode != nil {
		length++
		currentNode = currentNode.next
	}

	return length
}
