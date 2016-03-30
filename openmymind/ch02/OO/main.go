package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

type Saiyan1 struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {

	goku := Saiyan{"goku", 100}
	Super1(goku)
	fmt.Printf("goku's power = %d\n", goku.Power)
	goku1 := &Saiyan{"goku1", 200}
	Super2(goku1)
	fmt.Printf("goku1's power = %d\n", goku1.Power)
	Super3(goku1)
	fmt.Printf("goku1's power = %d\n", goku1.Power)

	goku.Super4()
	fmt.Printf("goku1's power = %d\n", goku.Power)

	goku2 := NewSaiyan("Foysal", 2500)
	fmt.Printf("%s has %d\n", goku2.Name, goku2.Power)

	goku3 := Saiyan1{"Foysal", 1, &Saiyan{
		"Abdul Aziz", 2,
	}}
	fmt.Printf("%s has %d\nFather is %s\n", goku3.Name, goku3.Power, goku3.Father.Name)

	goku4 := &Saiyan2{
		Person: &Person{"Goku"},
		Power:  9001,
	}

	goku4.Person.Introduce()
	goku4.Introduce()
}

func Super1(s Saiyan) {
	s.Power += 200
}

func Super2(s *Saiyan) {
	s.Power += 100
}

func Super3(s *Saiyan) {
	s = &Saiyan{"goku3", 400}
}

func (s *Saiyan) Super4() {
	s.Power += 55
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{Name: name, Power: power}
}

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan2 struct {
	*Person
	Power int
}

func (s *Saiyan2) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
