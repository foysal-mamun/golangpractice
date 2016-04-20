package main

import (
	"fmt"
)

func main() {
	//var P person
	//P.name = "Foysal"
	//P.age = 25

	//P := person{"Foysal", 32}
	//P := person{name: "Foysal", age: 31}
	P := struct {
		name string
		age  int
	}{"Numa", 1}

	fmt.Printf("The person's name is %s and age is %d\n\n", P.name, P.age)

	tom := person{"Tom", 18}
	bob := person{"bob", 25}
	pul := person{"pul", 43}

	tbO, tbD := Older(tom, bob)
	tpO, tpD := Older(tom, pul)
	bpO, bpD := Older(bob, pul)

	fmt.Printf("Of %s and %s, %s is older by %d Years\n", tom.name, bob.name, tbO.name, tbD)
	fmt.Printf("Of %s and %s, %s is older by %d Years\n", tom.name, pul.name, tpO.name, tpD)
	fmt.Printf("Of %s and %s, %s is older by %d Years\n\n", bob.name, pul.name, bpO.name, bpD)

	mark := Student{Human: Human{"Mark", 12, 20, "123"}, specialty: "CSE"}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)
	mark.weight += 5
	fmt.Println("His weight is ", mark.weight)

	mark.Skills = []string{"Golang", "Perl"}
	fmt.Println("His skills are ", mark.Skills)

	mark.int = 3
	fmt.Println("Her Preferred number is ", mark.int)

	emp := Employee{Human{"foysal", 31, 90, "123"}, "CSE", "987"}
	fmt.Printf("His phone number are %s and %s\n\n", emp.phone, emp.Human.phone)
}

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

type Human struct {
	name   string
	age    int
	weight int
	phone  string
}
type Student struct {
	Human
	Skills
	int
	specialty string
}
type Skills []string
type Employee struct {
	Human
	specialty string
	phone     string
}
