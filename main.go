package main

import (
	"fmt"

	"github.com/krahul2024/go-lib/lists/doubleList"
)

func main() {
	list := doubleList.NewList[int]()
	list.PushBack(89)
	list.PushFront(12)
	list.PopBack()
	list.PushFront(36)
	list.PushFront(123)
	list.PopFront()
	front, _ := list.Front()
	back, _ := list.Back()
	fmt.Println(front, back, list.Size())
}
