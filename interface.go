package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func sayHi(hasName HasName) {
	fmt.Println("Hai", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func Upss() interface{} {
	return "Bye Bye"
}

func main() {
	human := Person{"Boba"}
	sayHi(human)
	fmt.Println(Upss())
}
