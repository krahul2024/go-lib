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

type List[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

type iterator[T comparable] struct {
	List  *List[T]
	Ptr   *node[T]
	Index int
}

//-------------------------Iterator Operations------------------------------------------

func Iterator[T comparable](List *List[T]) *iterator[T] {
	return &iterator[T]{List, nil, -1}
}

func (it *iterator[T]) Begin() *iterator[T] {
	return &iterator[T]{List: it.List, Ptr: it.List.head.next, Index: 0}
}

func (it *iterator[T]) End() *iterator[T] {
	return &iterator[T]{List: it.List, Ptr: it.List.tail.prev, Index: it.List.size - 1}
}

func (it *iterator[T]) Next() (*iterator[T], error) {
	if it == nil || it.Ptr == nil || it.List == nil || it.List.head.next == it.List.tail {
		return nil, errors.New("can't iterate over invalid/empty List")
	}
	if it.Ptr.next == it.List.tail {
		it.Ptr = nil
		it.Index = -1
		return it, nil
	}
	it.Ptr = it.Ptr.next
	it.Index++
	return it, nil
}

func (it *iterator[T]) Prev() (*iterator[T], error) {
	if it == nil || it.Ptr == nil || it.List == nil || it.List.head.next == it.List.tail {
		return nil, errors.New("can't iterate over invalid/empty List")
	}
	if it.Ptr.prev == it.List.head {
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
	if it == nil || it.List == nil || it.List.head.next == it.List.tail || it.Ptr == nil { // in any of the cases the List seems to be empty
		return value, errors.New("can't get value for invalid/non-existent node")
	}
	return it.Ptr.value, nil
}

//-------------------------------------List Operations--------------------------------------

func NewList[T comparable]() *List[T] {
	head := &node[T]{}
	tail := &node[T]{}
	head.next = tail
	tail.prev = head
	return &List[T]{head: head, tail: tail, size: 0}
}

func (List *List[T]) PushFront(value T) {
	tempNode := &node[T]{value: value}
	List.head.next.prev = tempNode
	tempNode.next = List.head.next
	tempNode.prev = List.head
	List.head.next = tempNode
	List.size++
}

func (List *List[T]) PushBack(value T) {
	tempNode := &node[T]{value: value}
	List.tail.prev.next = tempNode
	tempNode.prev = List.tail.prev
	List.tail.prev = tempNode
	tempNode.next = List.tail
	List.size++
}

func (List *List[T]) PopBack() (T, error) {
	var value T
	if List.head.next == List.tail {
		return value, errors.New("can't delete from an empty List")
	}
	prevNode := List.tail.prev
	List.tail.prev.prev.next = List.tail
	List.tail.prev = List.tail.prev.prev
	value = prevNode.value
	prevNode = nil
	List.size--
	return value, nil
}

func (List *List[T]) PopFront() (T, error) {
	var value T
	if List.head.next == List.tail {
		return value, errors.New("can't delete from an empty List")
	}
	tempNode := List.head.next
	List.head.next.next.prev = List.head
	List.head.next = List.head.next.next
	value = tempNode.value
	tempNode = nil
	List.size--
	return value, nil
}

func (List *List[T]) Front() (T, error) {
	var value T
	if List.size == 0 || List.head.next == List.tail {
		return value, errors.New("the List is empty")
	}
	return List.head.next.value, nil
}

func (List *List[T]) Back() (T, error) {
	var value T
	if List.size == 0 || List.head.next == List.tail {
		return value, errors.New("the List is empty")
	}
	return List.tail.prev.value, nil
}

func (List *List[T]) Size() int {
	return List.size
}

func (List *List[T]) Empty() bool {
	return List.size == 0
}

func (List *List[T]) At(idx int) (T, *iterator[T], error) {
	var value T
	// in case of invalid index we return an iterator which points to nil with an error and zero-value
	if idx < 0 || List.size == 0 || idx >= List.size {
		return value, &iterator[T]{List: List, Ptr: nil, Index: -1}, errors.New("invalid index for candidate List")
	}
	it := Iterator[T](List)
	for it = it.Begin(); it.Index != idx; {
		it.Next()
	}
	return it.Ptr.value, it, nil
}

/*
	Add/Update/Delete at a given index/iterator
	1. Add : just add the node at given index, to make it a bit more flexible a List of values can be added after a given index
	2. Update : Update can be done for a single occurence or for all the occurences or for the given index
	3. Delete : Delete can be done for a single occurence or for all the occurences or for the given index
	4. Batch operations for all the actions, batch delete, batch add, batch update
*/

func (List *List[T]) AddAt(value T, index int) (*iterator[T], error) {
	if index == List.size {
		List.PushBack(value)
		it := Iterator[T](List)
		return it.End(), nil
	}
	_, it, err := List.At(index)
	if err != nil {
		return it, err
	}
	node := &node[T]{value: value}
	node.prev = it.Ptr.prev
	node.next = it.Ptr
	it.Ptr.prev.next = node
	it.Ptr.prev = node
	List.size++
	return it, nil
}

func (List *List[T]) RemoveAt(index int) (bool, error) {
	_, it, err := List.At(index)
	if err != nil {
		return false, err
	}
	it.Ptr.prev.next = it.Ptr.next
	it.Ptr.next.prev = it.Ptr.prev
	it.Ptr = nil
	List.size--
	return true, nil
}

func (List *List[T]) Remove(value T) (bool, error) {
	it := Iterator[T](List)
	for it = it.Begin(); it.Ptr != nil; {
		if value == it.Ptr.value {
			removed, err := List.RemoveAt(it.Index)
			return removed, err
		}
		it.Next()
	}
	return false, errors.New("element not found")
}

func (List *List[T]) RemoveAll(value T) (int, error) {
	var deletedCount = 0
	it := Iterator[T](List)
	for it = it.Begin(); it.Ptr != nil; {
		if value == it.Ptr.value {
			deleted, err := List.Remove(value)
			if err != nil {
				return deletedCount, err
			} else if deleted {
				deletedCount++
			}
		}
	}
	return deletedCount, nil
}

func (List *List[T]) Update(value T, index int) (*iterator[T], error) {
	_, it, err := List.At(index)
	if err != nil {
		return it, err
	}
	it.Ptr.value = value
	return it, nil
}

func (List *List[T]) Elements() []T {
	it := Iterator[T](List)
	var elements []T

	for it = it.Begin(); it.Ptr != nil; {
		elements = append(elements, it.Ptr.value)
		it.Next()
	}
	return elements
}

func (List *List[T]) Find(value T) (bool, *iterator[T], error) {
	it := Iterator[T](List)
	for it = it.Begin(); it.Ptr != nil; {
		if it.Ptr.value == value {
			return true, it, nil
		}
		it.Next()
	}
	return false, it, errors.New("element not found")
}

func (List *List[T]) FindAll(value T) []int {
	it := Iterator[T](List)
	var elements []int
	for it = it.Begin(); it.Ptr != nil; {
		if it.Ptr.value == value {
			elements = append(elements, it.Index)
		}
		it.Next()
	}
	return elements
}

func (List *List[T]) Has(value T) bool {
	has, _, _ := List.Find(value)
	return has
}
