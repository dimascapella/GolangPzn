package main

import "fmt"

type Odd func(int) bool

func main() {
	profile("rola", "Cap")
	fmt.Println(sum(10, 12))
	name, nickname, age := fullName()
	fmt.Println(name, nickname, age)
	nameSC, addressSC, oldSC := school()
	fmt.Println(nameSC, addressSC, oldSC)
	word := assemble("D", "i", "m", "a", "s")
	fmt.Println(word)
	// Variadic with slice
	para := []string{"a", "b", "c"}
	word2 := assemble(para...)
	fmt.Println(word2)
	olds := old
	fmt.Println(olds(10, 12))
	askingNumber(20, isOdd)
	fmt.Println(rec(5))
}

// function without params
func hewwo() {
	fmt.Println("Banana")
}

// Function with Params
func profile(firstname string, lastname string) {
	hewwo()
	fmt.Println(firstname, lastname)
}

// Function with Return
func sum(a int, b int) int {
	return a + b
}

// Function return multiple value
func fullName() (string, string, int) {
	return "Dimas", "Rola", 22
}

// Function return named value
func school() (name, address string, old int) {
	name = "SMA Negeri 1"
	address = "Jln. Gunung Jati"
	old = 65
	return name, address, old
}

// variadic function (input multiple data with single params)
func assemble(aplha ...string) string {
	var word string
	for _, value := range aplha {
		word += value
	}
	return word
}

// Function value
func old(a int, b int) int {
	return a + b
}

// Function as params
func isOdd(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

// func askingNumber(number int, check func(int) bool) {
// 	fmt.Println("Number ", number, "is ", check(number))
// }

// Function as params with type dec
func askingNumber(number int, odd Odd) {
	fmt.Println("Number ", number, "is ", odd(number))
}

// Recursive Function
func rec(number int) int {
	if number <= 1 {
		return number
	} else {
		return number * rec(number-1)
	}
}
