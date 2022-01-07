package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushFront("Dimas")
	data1 := data.PushFront("Eka")
	data.PushFront("Adinandra")
	data.PushBack("DEA")
	data.InsertAfter("Hallo", data1)
	data.InsertBefore("Bleh", data1)

	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
