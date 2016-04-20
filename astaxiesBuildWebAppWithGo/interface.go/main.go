package main

import (
	"fmt"
)

func name() {

}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mike := Student{Human{"mike", 25, "222"}, "MIT", 0.00}
	paul := Student{Human{"paul", 26, "111"}, "Harvard", 100}
	sam := Employee{Human{"sam", 36, "444"}, "Golang Inc", 1000}
	tom := Employee{Human{"tom", 25, "xxx"}, "Things Ltd", 5000}

	var i Men

	i = mike
	fmt.Println("This is Mike, a Student")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 4)
	x[0], x[1], x[2], x[3] = paul, sam, mike, tom

	for _, value := range x {
		value.SayHi()
	}
}
