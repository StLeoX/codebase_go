package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string `label:"name is "`
	Age    int    `label:"age is "`
	Gender string `label:"gender is " default:"unknown"`
}

func (person Person) Info() {
	fmt.Printf("%s,%d,%s\n", person.Name, person.Age, person.Gender)
}

func Print(obj interface{}) error {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		// get
		field := v.Type().Field(i)
		tag := field.Tag
		// parse
		label := tag.Get("label")
		defaultV := tag.Get("default")
		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			value = defaultV
		}
		fmt.Println(label + value + v.String())
	}
	return nil
}

func main() {
	person := Person{
		Name:   "1",
		Age:    0,
		Gender: "",
	}
	person.Info()
	fmt.Printf("====")
	Print(person)
}
