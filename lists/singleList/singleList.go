package singleList

import (
	"errors"
)

//------------------------Type Declarations---------------------------------

type list[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

type node[T comparable] struct {
	value T
	next  *node[T]
}

type iterator[T comparable] struct {
	*node[T]
}

//------------------------Iterator Operations------------------------------------

func (list *list[T]) Iterator() *iterator[T] {
	if list.size == 0 {
		return nil
	}
	return &iterator[T]{list.head}
}

// func (it *iterator[T]) End() bool {
// 	return it.current == nil
// }

func (it *iterator[T]) Value() T {
	var zeroValue T
	if it == nil {
		return zeroValue
	}
	return it.value
}

func (it *iterator[T]) Next() *iterator[T] {
	if it.next == nil {
		return nil
	}
	return &iterator[T]{it.next}
}

//--------------------Linked List Operations--------------------------------------

func NewList[T comparable](values ...T) *list[T] {
	list := &list[T]{}
	for _, value := range values {
		list.PushBack(value)
	}
	return list
}

func (list *list[T]) Front() T {
	var zeroValue T
	if list.head == nil {
		return zeroValue
	}
	return list.head.value
}

func (list *list[T]) Back() T {
	var zeroValue T
	if list.tail == nil {
		return zeroValue
	}
	return list.tail.value
}

func (list *list[T]) PushBack(value T) {
	tempNode := &node[T]{value: value}
	if list.head == nil {
		list.head = tempNode
		list.tail = tempNode
	} else {
		list.tail.next = tempNode
		list.tail = tempNode
	}
	list.size++
}

func (list *list[T]) PushFront(value T) {
	node := &node[T]{value: value}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.size++
}

func (list *list[T]) Size() int {
	return list.size
}

func (list *list[T]) PopFront() error {
	if list.head == nil {
		return errors.New("can't delete from empty list")
	}
	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size = 0
	} else {
		list.head = list.head.next
		list.size--
	}
	return nil
}

func (list *list[T]) PopBack() error {
	if list.head == nil {
		return errors.New("can't delete from empty list")
	}
	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size = 0
		return nil
	} else {
		tempNode := list.head
		for tempNode.next.next != nil {
			tempNode = tempNode.next
		}
		tempNode.next = nil
		list.tail = tempNode
		list.size--
	}
	return nil
}

func (list *list[T]) Remove(itr *iterator[T]) error {
	if itr == nil {
		return errors.New("nothing to delete")
	}
	// if the iterator points to nil, it a list with single element or the iterator is at the end of list
	if itr.next == nil {
		// list with single value
		if itr.next == list.head.next {
			itr = nil
			return list.PopFront()
		}
		itr = nil
		return list.PopBack() // iterator is at the last element of list
	}
	tempNode := list.head
	for tempNode.next != itr.next {
		tempNode = tempNode.next
	}
	if tempNode != nil {
		tempNode.value = tempNode.next.value
		tempNode.next = tempNode.next.next
		itr = nil
		list.size--
		return nil
	} else { // iterator points to invalid address
		itr = nil
		return nil
	}
}

func (list *list[T]) Find(val T) (int, error) {
	index := -1
	tempNode := list.head
	for tempNode != nil {
		index++
		if tempNode.value == val {
			return index, nil
		}
		tempNode = tempNode.next
	}
	return -1, errors.New("element not found in the list")
}

func (list *list[T]) Value(it *iterator[T]) (T, error) {
	var zeroValue T
	if it == nil {
		return zeroValue, errors.New("invalid iterator")
	}
	return it.value, nil
}

func (list *list[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *list[T]) Empty() bool {
	return list.size == 0
}
