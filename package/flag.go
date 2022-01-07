package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put yout database host")
	user := flag.String("user", "root", "Put yout database user")
	password := flag.String("password", "root", "Put yout database password")
	number := flag.Int("number", 100, "Put yout number")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*password)
	fmt.Println(*number)
}
