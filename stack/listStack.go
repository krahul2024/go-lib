package stack

import (
	"errors"

	"github.com/krahul2024/go-lib/lists/doubleList"
)

//--------------------------- Type Declarations-------------------------
type Stack[T comparable] struct {
	list *doubleList.List[T]
}

//----------------------Stack Operations-----------------------------------

func NewStack[T comparable]() *Stack[T] {
	list := doubleList.NewList[T]()
	return &Stack[T]{list}
}

func (stack *Stack[T]) Push(value T) int {
	stack.list.PushFront(value)
	return stack.list.Size()
}

func (stack *Stack[T]) Top() (T, error) {
	var value T
	if stack.list.Size() == 0 {
		return value, errors.New("empty stack")
	}
	return stack.list.Front()
}

func (stack *Stack[T]) Pop() (T, error) {
	var value T
	if stack.list.Size() == 0 {
		return value, errors.New("can't pop empty stack")
	}
	return stack.list.PopFront()
}

func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

func (stack *Stack[T]) Has(value T) bool {
	return stack.list.Has(value)
}

func (stack *Stack[T]) Empty() bool {
	return stack.list.Size() == 0
}

// func (stack *Stack[T]) Clear() bool {

// }
