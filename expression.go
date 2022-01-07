package main

import "fmt"

func main() {
	// if else statement
	name := "Dimas"

	if name == "Dimas" {
		fmt.Println("Bleh")
	} else if name == "Eka" {
		fmt.Println("Bleh Bleh")
	} else {
		fmt.Println("Bleh Bleh Bleh")
	}

	// Short Statement if else
	if length := len(name); length > 3 {
		fmt.Println("Bleh Bleh")
	} else {
		fmt.Println("Banana")
	}

	// Switch Expression
	switch name {
	case "Dimas":
		fmt.Println("Bleh")
	case "Eka":
		fmt.Println("Bleh Bleh")
	default:
		fmt.Println("Bleh Bleh Bleh")
	}

	// short statement switch expression
	switch length := len(name); length < 3 {
	case true:
		fmt.Println("Bleh Bleh")
	case false:
		fmt.Println("Banana")
	}

	// switch expression without condition
	length2 := len(name)
	switch {
	case length2 > 5:
		fmt.Println("Bleh")
	case length2 < 1:
		fmt.Println("Bleh Bleh")
	default:
		fmt.Println("Bleh Bleh Bleh")
	}

	// For Loops Expression
	// Option 1
	counter := 1
	for counter <= 10 {
		fmt.Println("Counting = ", counter)
		counter++
	}

	// Option 2
	for i := 1; i <= 10; i++ {
		fmt.Println("Loopy Woopy = ", i)
	}

	// Option 3
	for _, j := range name {
		fmt.Println(j)
	}
}
