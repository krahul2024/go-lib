package main

import (
	"fmt"

	"github.com/krahul2024/go-lib/lists/doubleList"
)

func main() {

	dlist := doubleList.NewList(1)
	dlist.PushFront(12)
	dlist.PushBack(31)
	dlist.PushFront(600)
	dlist.PushBack(90)
	dlist.PushBack(81)
	it := dlist.Iterator()
	for it != nil {
		fmt.Print(it.Value(), " ") // Print the value of the node the iterator is pointing to
		it = it.Next()
	}
	fmt.Println(dlist.Size())

	dlist.PopBack()
	it = dlist.Iterator()
	dlist.PopFront()
	dlist.PopFront()
	for it != nil {
		fmt.Print(it.Value(), " ") // Print the value of the node the iterator is pointing to
		it = it.Next()
	}
	fmt.Println()
	it = dlist.Iterator()
	dlist.PushBack(301)
	dlist.PushFront(60000)
	for it != nil {
		fmt.Print(it.Value(), " ") // Print the value of the node the iterator is pointing to
		it = it.Next()
	}
	fmt.Println(dlist.Size())
}
