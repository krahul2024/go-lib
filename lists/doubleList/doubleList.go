package doubleList

import (
	"errors"
)

//-------------------------Type Declarations-----------------------------------------------

type node[T comparable] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

type list[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

type iterator[T comparable] struct {
	*node[T]
}

//-------------------------Iterator Operations------------------------------------------

func (list *list[T]) Begin() *iterator[T] {

}

//-------------------------------------List Operations--------------------------------------

func NewList[T comparable]() *list[T] {
	head := &node[T]{}
	tail := &node[T]{}
	head.next = tail
	tail.prev = head
	return &list[T]{head: head, tail: tail, size: 0}
}

func (list *list[T]) PushFront(value T) {
	tempNode := &node[T]{value: value}
	list.head.next.prev = tempNode
	tempNode.next = list.head.next
	tempNode.prev = list.head
	list.head.next = tempNode
	list.size++
}

func (list *list[T]) PushBack(value T) {
	tempNode := &node[T]{value: value}
	list.tail.prev.next = tempNode
	tempNode.prev = list.tail.prev
	list.tail.prev = tempNode
	tempNode.next = list.tail
	list.size++
}

func (list *list[T]) PopBack() (T, error) {
	var value T
	if list.head.next == list.tail {
		return value, errors.New("can't delete from an empty list")
	}
	prevNode := list.tail.prev
	list.tail.prev.prev.next = list.tail
	list.tail.prev = list.tail.prev.prev
	value = prevNode.value
	prevNode = nil
	list.size--
	return value, nil
}

func (list *list[T]) PopFront() (T, error) {
	var value T
	if list.head.next == list.tail {
		return value, errors.New("can't delete from an empty list")
	}
	tempNode := list.head.next
	list.head.next.next.prev = list.head
	list.head.next = list.head.next.next
	value = tempNode.value
	tempNode = nil
	list.size--
	return value, nil
}

func (list *list[T]) Front() (T, error) {
	var value T
	if list.size == 0 || list.head.next == list.tail {
		return value, errors.New("the list is empty")
	}
	return list.head.next.value, nil
}

func (list *list[T]) Back() (T, error) {
	var value T
	if list.size == 0 || list.head.next == list.tail {
		return value, errors.New("the list is empty")
	}
	return list.tail.prev.value, nil
}

func (list *list[T]) Size() int {
	return list.size
}

func (list *list[T]) Empty() bool {
	return list.size == 0
}

// h[nil, t] <-> t[h, nil]
