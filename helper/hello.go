package helper

import "fmt"

// access modifier public (Capitalize)
const Application = "Learning Golang"

// access modifier private (lowercase)
const version = "1"

func Hi(Name string) {
	fmt.Println("Hewwo", Name)
}
