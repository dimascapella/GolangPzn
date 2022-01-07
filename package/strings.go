package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Dimas Eka Adinandra", "a"))
	fmt.Println(strings.Split("Dimas Eka Adinandra", " "))
	fmt.Println(strings.ToLower("Dimas Eka Adinandra"))
	fmt.Println(strings.ToUpper("Dimas Eka Adinandra"))
	fmt.Println(strings.ToTitle("dimas eka adinandra"))
	fmt.Println(strings.Trim("wdimas eka adinandraw", "w"))
	fmt.Println(strings.ReplaceAll("Dimas Eka Adinandra", "Eka", "Ake"))
}
