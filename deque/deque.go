package deque

import (
	"errors"

	"github.com/krahul2024/go-lib/lists/doubleList"
)

//-------------------------Declarations-----------------------------

type Deque[T comparable] struct {
	list *doubleList.List[T]
}

//----------------------------------Deque Operations--------------------------------
func NewDeque[T comparable]() *Deque[T] {
	list := doubleList.NewList[T]()
	return &Deque[T]{list}
}

func (d *Deque[T]) PushFront(value T) int {
	d.list.PushFront(value)
	return d.list.Size()
}

func (d *Deque[T]) PushBack(value T) int {
	d.list.PushBack(value)
	return d.list.Size()
}

func (d *Deque[T]) Front() (T, error) {
	var value T
	if d.list.Size() == 0 {
		return value, errors.New("queue is empty")
	}
	return d.list.Front()
}

func (d *Deque[T]) Back() (T, error) {
	var value T
	if d.list.Size() == 0 {
		return value, errors.New("queue is empty")
	}
	return d.list.Back()
}

func (d *Deque[T]) PopFront() (T, error) {
	var value T
	if d.list.Size() == 0 {
		return value, errors.New("empty queue, nothing to pop")
	}
	return d.list.PopFront()
}

func (d *Deque[T]) PopBack() (T, error) {
	var value T
	if d.list.Size() == 0 {
		return value, errors.New("empty queue, nothing to pop")
	}
	return d.list.PopBack()
}

func (d *Deque[T]) Size() int {
	return d.list.Size()
}

func (d *Deque[T]) Empty() bool {
	return d.list.Size() == 0
}

// add clear function, check out the time complexity of that
