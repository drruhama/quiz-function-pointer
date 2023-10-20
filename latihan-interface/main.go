package main

import "fmt"

type Calculate interface {
	Substract() float32
}

type Num1 struct {
	Num         float32
	Substracter float32
}

func NewNum1(num, substracter float32) Calculate {
	return &Num1{Num: num, Substracter: substracter}
}

func (n Num1) Substract() float32 {
	return n.Num - n.Substracter
}

func main() {
	num := NewNum1(4.0, 2.0)
	fmt.Println(num.Substract())
}
