package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s Student) SetMyName(nama *string, name string) {
	*nama = name
}

func (s Student) CallMyName() {
	fmt.Println("Hello my name is :", s.Name)
}

func main() {
	student := Student{
		Name:  "Nobee",
		Class: "1A",
	}
	myName := &student.Name
	student.SetMyName(myName, "David")
	student.CallMyName()
}
