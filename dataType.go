package main

import "fmt"

func main() {

	// // Data Type Int and Float (Number)
	// fmt.Println(int8(125))
	// fmt.Println(int16(10000))
	// fmt.Println(int32(1000000000))
	// fmt.Println(int64(1000000000000000000))
	// fmt.Println(float32(3.141231231))
	// fmt.Println(float64(3.141231231))

	// // Data Type Boolean
	// fmt.Println(true)
	// fmt.Println(false)

	// // Data Type String
	// fmt.Println("Learning Golang")
	// // String Function
	// // >> Len("String") Counting length of Word
	// fmt.Println(len("Learning Golang"))
	// // >> "String"[index] Get a value from word with index
	// fmt.Println("Learning Golang"[3])

	// Data Type Array
	// Option 1
	// var names [3]string
	// names[0] = "Dimas"
	// names[1] = "Eka"
	// names[2] = "Adinandra"

	// fmt.Println(names)
	// // Option 2
	// animals := [3]string{
	// 	"Dog",
	// 	"Cat",
	// 	"Dojjo",
	// }

	// fmt.Println(animals)
	// animals[1] = "Cats"
	// fmt.Println(animals)

	// // Data Type Slice
	// monts := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// slice := monts[3:7]
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// // Capacity Func
	// fmt.Println(cap(slice))
	// // Append Func
	// slice = append(slice, 101)
	// fmt.Println(slice)
	// // Make Func
	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "Ara"
	// newSlice[1] = "Rola"
	// fmt.Println(newSlice)
	// // Copy Func
	// copySlice := make([]string, len(newSlice), cap(newSlice))
	// copy(copySlice, newSlice)
	// fmt.Println(copySlice)

	// Data Type Map
	bio := map[string]string{
		"name":    "Dimas",
		"Address": "Gunung Jati",
	}
	fmt.Println(bio["name"])

	books := make(map[string]string)
	books["Title"] = "The Lord of BDO"
	books["Wrong"] = "Mueehehheh"
	books["Class"] = "Mystic"

	delete(books, "Wrong")
	fmt.Println(books)
}
