package main

import "fmt"

func error() interface{} {
	return "Waw"
}

func main() {
	temp := error()
	switch res := temp.(type) {
	case string:
		fmt.Println("String Type", res)
	case int:
		fmt.Println("Int Type", res)
	default:
		fmt.Println("Unknown", res)
	}
}
