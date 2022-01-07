package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("d([a-zA-Z0-9])s")

	fmt.Println(regex.MatchString("dis"))
	fmt.Println(regex.MatchString("d0s"))

	search := regex.FindAllString("dio dis d8s dar", -1)
	fmt.Println(search)
}
