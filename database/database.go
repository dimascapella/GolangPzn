package database

import "fmt"

var connection string

func init() {
	connection = "MySQL"
	fmt.Println("Running Package")
}

func GetConnection() string {
	return connection
}
