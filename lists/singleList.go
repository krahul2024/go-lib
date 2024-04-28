package lists

type list[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

type iterator[T comparable] struct {
	current *node[T]
}

type node[T comparable] struct {
	value T
	next  *node[T]
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

func NewList[T comparable](values ...T) *list[T] {
	list := &list[T]{}
	for _, value := range values {
		list.PushBack(value)
	}
	return list
}

func (list *list[T]) Iterator() *iterator[T] {
	return &iterator[T]{list.head}
}

func (it *iterator[T]) End() bool {
	return it.current == nil
}

func (it *iterator[T]) Value() T {
	var zeroValue T
	if it.current == nil {
		return zeroValue
	}
	return it.current.value
}
