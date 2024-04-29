package main

import (
	"fmt"

	"github.com/krahul2024/go-lib/lists"
)

func main() {

	list := lists.NewList[string]("first", "second", "third", "fourth")
	// list.PushBack("fourth")
	// list.PushFront("zero'th")

	it := list.Iterator()
	count := 0
	for it != nil && count < 2 {
		fmt.Println(it.Value())
		it = it.Next()
		count++
	}
	fmt.Println(it.Value(), it)
	index, err := list.Find("fourth")
	result, err := list.Value(it)
	fmt.Println(index, err, result)
}
