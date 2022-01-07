package main

import "fmt"

func main() {
	// Writing Variable (var)
	// Option 1
	var name string
	name = "Dimas Eka Adinandra"
	fmt.Println(name)

	// Option 2
	var name1 = "Dimas E.A"
	fmt.Println(name1)

	// Option 3
	name2 := "Dimas Eka"
	fmt.Println(name2)

	// Multiple Variable Assignment
	var (
		firstname = "Dimas"
		lastname  = "Eka A"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)

	// Writing Variable (const)
	// const nm string
	// nm = "Rola"

	// const nm1 = "Rola"
	// nm2 := "Rola"

	// const (
	// 	nm3 = "Rola"
	// 	nm4 = 100
	// )
}
