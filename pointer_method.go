package main

import "fmt"

type Human struct {
	Name string
}

func (human *Human) Married() {
	human.Name = "Mr. " + human.Name
}

func main() {
	robot := Human{"Robo"}
	robot.Married()
	fmt.Println(robot.Name)
}
