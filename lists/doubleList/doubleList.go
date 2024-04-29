package doubleList

import (
	"errors"
)

//----------------Type Declarations--------------------------

type dlist[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

type node[T comparable] struct {
	next  *node[T]
	prev  *node[T]
	value T
}

type iterator[T comparable] struct {
	*node[T]
}

//-----------------------Iterator Operations----------------------------

func (list *dlist[T]) Iterator() *iterator[T] {
	if list.head == nil {
		return nil
	}
	return &iterator[T]{list.head}
}

func (it *iterator[T]) Value() T {
	var zeroValue T
	if it == nil {
		return zeroValue
	}
	return it.value
}

func (it *iterator[T]) Next() *iterator[T] {
	if it == nil || it.next == nil { // doing it this way instead of just returning it.next when it != nil is because it.next is a struct pointer and this would result in nil value like this &{nil}
		return nil
	}
	return &iterator[T]{it.next}
}

func (it *iterator[T]) Prev() *iterator[T] {
	if it == nil || it.prev == nil {
		return nil
	}
	return &iterator[T]{it.prev}
}

//-------------------List Operations-------------------------------------

func NewList[T comparable](values ...T) *dlist[T] {
	dlist := &dlist[T]{}
	for _, value := range values {
		dlist.PushBack(value)
	}
	return dlist
}

func (list *dlist[T]) PushBack(value T) {
	tempNode := &node[T]{value: value}

	if list.head == nil {
		list.head = tempNode
		list.tail = tempNode
	} else {
		tempNode.prev = list.tail
		list.tail.next = tempNode
		list.tail = tempNode
	}
	list.size++
}

func (list *dlist[T]) Size() int {
	return list.size
}

func (list *dlist[T]) Front() (T, error) {
	var zeroValue T
	if list.head == nil {
		return zeroValue, errors.New("can't peek an empty list")
	}
	return list.head.value, nil
}

func (list *dlist[T]) Back() (T, error) {
	var zeroValue T
	if list.head == nil {
		return zeroValue, errors.New("no element at the back for an empty list")
	}
	return list.tail.value, nil
}
