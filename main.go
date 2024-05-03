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
	list2.PushBack("first")
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

	// _, itr, _ := list2.At(0)
	// fmt.Println(itr.Value())
	// itr.Next()
	// fmt.Println(itr.Value())
	// list2.Update("updated value", itr.Index)
	// fmt.Println(itr.Value())
	// itr.Next()
	// fmt.Println(itr.Value())

	// for itr = itr.Begin(); itr.Ptr != nil; {
	// 	fmt.Println(itr.Value())
	// 	itr.Next()
	// }

	fmt.Println(list2.FindAll("second"))
	fmt.Println(list2.FindAll("first"))
	fmt.Println(list2.FindAll("fasdfsaf"))

	list2.AddAt("value_at_zeroth index", 0)
	fmt.Println(list2.Elements())
	list2.AddAt("value_at_zeroth index", -1)
	fmt.Println(list2.Elements())
	list2.AddAt("value_at_zeroth index", list2.Size()+1)
	fmt.Println(list2.Elements())
	list2.AddAt("value_at_zeroth index", list2.Size())
	fmt.Println(list2.Elements())
	list2.AddAt("value_at_second_index", 2)
	fmt.Println(list2.Elements(), list2.Size())

	list2.RemoveAt(0)
	fmt.Println(list2.Elements(), list2.Size())

	list2.Remove("first")
	fmt.Println(list2.Elements(), list2.Size())

}
