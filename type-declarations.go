package main

import "fmt"

func main() {
	type name string
	type graduation bool

	const nm name = "Rola"
	const gd graduation = false

	fmt.Println(nm, gd)
}
