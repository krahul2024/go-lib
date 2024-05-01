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
	list  *list[T]
	Ptr   *node[T]
	Index int
}

//-------------------------Iterator Operations------------------------------------------

func Iterator[T comparable](list *list[T]) *iterator[T] {
	return &iterator[T]{list, nil, -1}
}

func (it *iterator[T]) Begin() *iterator[T] {
	return &iterator[T]{list: it.list, Ptr: it.list.head.next, Index: 0}
}

func (it *iterator[T]) End() *iterator[T] {
	return &iterator[T]{list: it.list, Ptr: it.list.tail.prev, Index: it.list.size - 1}
}

func (it *iterator[T]) Next() (*iterator[T], error) {
	if it == nil || it.Ptr == nil || it.list == nil || it.list.head.next == it.list.tail {
		return nil, errors.New("can't iterate over invalid/empty list")
	}
	if it.Ptr.next == it.list.tail {
		it.Ptr = nil
		it.Index = -1
		return it, nil
	}
	it.Ptr = it.Ptr.next
	it.Index++
	return it, nil
}

func (it *iterator[T]) Prev() (*iterator[T], error) {
	if it == nil || it.Ptr == nil || it.list == nil || it.list.head.next == it.list.tail {
		return nil, errors.New("can't iterate over invalid/empty list")
	}
	if it.Ptr.prev == it.list.head {
		it.Ptr = nil
		it.Index = -1
		return it, nil
	}
	it.Ptr = it.Ptr.prev
	it.Index--
	return it, nil
}

func (it *iterator[T]) Value() (T, error) {
	var value T
	if it == nil || it.list == nil || it.list.head.next == it.list.tail || it.Ptr == nil { // in any of the cases the list seems to be empty
		return value, errors.New("can't get value for invalid/non-existent node")
	}
	return it.Ptr.value, nil
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

func (list *list[T]) At(idx int) (T, *iterator[T], error) {
	var value T
	// in case of invalid index we return an iterator which points to nil with an error and zero-value
	if idx < 0 || list.size == 0 || idx >= list.size {
		return value, &iterator[T]{list: list, Ptr: nil, Index: -1}, errors.New("invalid index for candidate list")
	}
	it := Iterator[T](list)
	for it = it.Begin(); it.Index != idx; {
		it.Next()
	}
	return it.Ptr.value, it, nil
}

/*
	Add/Update/Delete at a given index/iterator
	1. Add : just add the node at given index, to make it a bit more flexible a list of values can be added after a given index
	2. Update : Update can be done for a single occurence or for all the occurences or for the given index
	3. Delete : Delete can be done for a single occurence or for all the occurences or for the given index
	4. Batch operations for all the actions, batch delete, batch add, batch update
*/

func (list *list[T]) Update(value T, index int) (*iterator[T], error) {
	_, it, err := list.At(index)
	if err != nil {
		return it, err
	}
	it.Ptr.value = value
	return it, nil
}

func (list *list[T]) Elements() []T {
	it := Iterator[T](list)
	var elements []T

	for it = it.Begin(); it.Ptr != nil; {
		elements = append(elements, it.Ptr.value)
		it.Next()
	}
	return elements
}

func (list *list[T]) Find(value T) (bool, *iterator[T], error) {
	it := Iterator[T](list)
	for it = it.Begin(); it.Ptr != nil; {
		if it.Ptr.value == value {
			return true, it, nil
		}
		it.Next()
	}
	return false, it, errors.New("element not found")
}

func (list *list[T]) FindAll(value T) []int {
	it := Iterator[T](list)
	var elements []int
	for it = it.Begin(); it.Ptr != nil; {
		if it.Ptr.value == value {
			elements = append(elements, it.Index)
		}
		it.Next()
	}
	return elements
}

func (list *list[T]) Has(value T) bool {
	has, _, _ := list.Find(value)
	return has
}
