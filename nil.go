package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	hela := newMap("Boba")
	if hela == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(hela)
	}
}
