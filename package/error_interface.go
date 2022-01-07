package main

import (
	"errors"
	"fmt"
)

func divide(number, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("No Zero Divier")
	} else {
		result := number / divider
		return result, nil
	}
}

func main() {
	result, error := divide(10, 5)
	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}
}
