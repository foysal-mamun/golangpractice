package main

import (
	"fmt"
	"math"
)

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())

	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

	boxex := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxex))
	fmt.Println("The value of the first one is: ", boxex[0].Valume(), "cm3")
	fmt.Println("The color of the last one is: ", boxex[len(boxex)-1].color.String())
	fmt.Println("The biggest one is ", boxex.BiggestsColor().String())

	fmt.Println("Let's paint them all black")
	boxex.PaintItBlack()
	fmt.Println("The color of the second one is ", boxex[0].color.String())

	fmt.Println("Obviously, now, the biggets one is ", boxex.BiggestsColor().String())

	mark := Student{Human{"mark", 25, "123"}, "MIT"}
	sarm := Employee{Human{"sam", 45, "111"}, "Golang Inc"}

	mark.SayHi()
	sarm.SayHi()

}

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s\n", e.name, e.company, e.phone)
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c *Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box

func (b Box) Valume() float64 {
	return b.width * b.height * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}
func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)

	for _, b := range bl {
		if bv := b.Valume(); bv > v {
			v = bv
			k = b.color
		}
	}

	return k
}
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
