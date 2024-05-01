package main

import (
	"fmt"

	"github.com/krahul2024/go-lib/lists/doubleList"
)

func main() {
	// list := doubleList.NewList[int]()
	list2 := doubleList.NewList[string]()
	// it := doubleList.Iterator[int](list)
	// it2 := doubleList.Iterator[string](list2)

	// list.PushBack(122)
	// list.PushBack(153)
	// list.PushFront(84)
	// list.PushBack(312)

	// for it = it.Begin(); it.Ptr != nil; {
	// 	val, _ := it.Value()
	// 	fmt.Println(val, it.Index, list.Size())
	// 	it.Next()
	// }

	// list.PushFront(98)
	// list.PopBack()

	// for it = it.End(); it.Ptr != nil; {
	// 	val, _ := it.Value()
	// 	fmt.Println(val, it.Index, list.Size())
	// 	it.Prev()
	// }

	list2.PushBack("first")
	list2.PushFront("zero'th")
	list2.PushBack("second")
	// for it2 = it2.Begin(); it2.Ptr != nil; {
	// 	fmt.Println(it2.Value())
	// 	it2.Next()
	// }

	// for count := -2; count < 6; count++ {
	// 	value, itr, err := list2.At(count)
	// 	fmt.Println(count, "value : ", value, "Iterator : ", itr, "Error : ", err)
	// 	itr, err = itr.Next()
	// 	fmt.Println(err)
	// 	value, err = itr.Value()
	// 	fmt.Println("Next Value : ", value, err)
	// 	fmt.Println()
	// }

	_ , itr, _ := list2.At(1)
	fmt.Println(itr.Value())
	itr.Next()
	fmt.Println(itr.Value())

}
