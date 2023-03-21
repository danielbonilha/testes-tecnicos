package datastructures

import (
	"errors"
	"fmt"
)

type (
	LinkedList struct {
		head *item
		tail *item
		size int
	}

	item struct {
		data     any
		previous *item
		next     *item
	}
)

// NewLinkedList creates a new LinkedList
// using the provided data at head node
// Costs O(1)
func NewLinkedList(data any) LinkedList {
	item := &item{
		data:     data,
		previous: nil,
		next:     nil,
	}

	return LinkedList{
		size: 1,
		head: item,
		tail: item,
	}
}

// InsertHead insert new data to the head
// Costs O(1)
func (l *LinkedList) InsertHead(data any) {
	// create an item with the data;
	// points to the head as next item
	// have no previous item
	newLinkedListItem := &item{
		data:     data,
		previous: nil,
		next:     l.head,
	}
	// assign the previous node of the head to the new item
	if l.head != nil {
		l.head.previous = newLinkedListItem
	}
	// assign the head to the new item
	l.head = newLinkedListItem
	l.size += 1
}

// InsertTail insert new data to the tail
// Costs O(1)
func (l *LinkedList) InsertTail(data any) {
	// create an item with the data
	// that points to the tail as previous item
	// and that has no next item
	newLinkedListItem := &item{
		data:     data,
		previous: l.tail,
		next:     nil,
	}
	// assign the next node of the tail to the new item
	if l.tail != nil {
		l.tail.next = newLinkedListItem
	}
	// assign the head to the new item
	l.tail = newLinkedListItem
	l.size += 1
}

// DeleteHead changes the head reference to the
// second element of the list
// Costs O(1)
func (l *LinkedList) DeleteHead() error {
	if l.size == 0 {
		return errors.New("stack underflow")
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return nil
	}

	l.head = l.head.next
	l.head.previous = nil
	l.size -= 1

	return nil
}

func (l *LinkedList) DeleteTail() error {
	if l.size == 0 {
		return errors.New("stack underflow")
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return nil
	}

	l.tail = l.tail.previous
	l.tail.next = nil
	l.size -= 1

	return nil
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Display() {
	node := l.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}
