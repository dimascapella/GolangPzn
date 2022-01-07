package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	school        bool
}

// Pointer Function
func changeAddress(customer *Customer) {
	customer.Address = "HUBLAHHHHHHH"
}

func main() {
	// // Pass By Value
	// customer1 := Customer{"Boba", "Gudu", 30, false}
	// customer2 := customer1

	// customer2.Address = "Bleh"

	// fmt.Println(customer1)
	// fmt.Println(customer2)

	// // Pointer (Mengubah Data Struct Dengan key)
	// customer1 := Customer{"Boba", "Gudu", 30, false}
	// customer2 := &customer1

	// customer2.Address = "Bleh"

	// fmt.Println(customer1)
	// fmt.Println(customer2)

	// Pointer (Mengubah Semua Data di Struct)
	customer1 := Customer{"Boba", "Gudu", 30, false}
	customer2 := &customer1

	customer2.Address = "Bleh"

	// Mengganti Data customer 2 ke data baru
	// customer2 = &Customer{"Lola", "Blewa", 30, false}
	// Mengganti Semua Pointer ke data baru
	*customer2 = Customer{"Lola", "Blewa", 30, false}

	fmt.Println(customer1)
	fmt.Println(customer2)

	// Pointer dengan menggunakan function new
	customer3 := new(Customer)
	customer3.Name = "Hublah"
	fmt.Println(customer3)

	customer4 := Customer{"Mana", "Aha", 15, true}
	changeAddress(&customer4)
	fmt.Println(customer4)
}
