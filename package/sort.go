package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].age < userSlice[j].age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Boba", 21},
		{"Allo", 51},
		{"Mona", 11},
		{"udi", 31},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
	fmt.Println(len(users))
}
