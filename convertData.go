package main

import "fmt"

func main() {
	// Convert Data Int
	int32 := 25913
	int8 := int8(int32)
	int16 := int16(int32)
	int64 := int64(int32)

	fmt.Println(int8)
	fmt.Println(int16)
	fmt.Println(int32)
	fmt.Println(int64)

	// Convert Data String
	name := "Dimas Eka Adinandra"
	index3 := string(name[3])
	fmt.Println(index3)
}
