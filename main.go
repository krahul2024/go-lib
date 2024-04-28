package main

import (
	"fmt"

	"github.com/krahul2024/go-lib/lists"
)

func main() {
	list := lists.NewList[int]()
	list.PushBack(13)
	it := list.Iterator()
	fmt.Println(it.Value())
}
