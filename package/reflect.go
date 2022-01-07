package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isRequired(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Dimas"}
	sampleType := reflect.TypeOf(sample)
	stuctField := sampleType.Field(0)
	required := stuctField.Tag.Get("required")
	fmt.Println(required)

	fmt.Println(isRequired(sample))
}
