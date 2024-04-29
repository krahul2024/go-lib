package main

import (
	"fmt"

	"github.com/krahul2024/go-lib/lists/doubleList"
)

func main() {

	dlist := doubleList.NewList(1, 110)
	dlist.PushBack(31)
	dlist.PushBack(90)
	dlist.PushBack(81)
	it := dlist.Iterator()
	for it != nil {
		fmt.Println(it.Value()) // Print the value of the node the iterator is pointing to
		it = it.Next()
	}

}
