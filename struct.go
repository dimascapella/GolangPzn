package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	// option 1
	var dimas Customer
	dimas.Name = "Dimas"
	dimas.Address = "Gunung Jati"
	dimas.Age = 22

	fmt.Println(dimas)

	// option 2
	human := Customer{
		Name:    "Humanoid",
		Address: "Mars",
		Age:     1050,
	}
	fmt.Println(human)

	// option 3
	human2 := Customer{"Angke", "Pluto", 2000}
	fmt.Println(human2)
	human2.sayHai("Bambang")
}
