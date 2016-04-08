package main

import (
	"fmt"
)

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func main() {

	lookup := make(map[string]int)
	lookup["foysal"] = 100
	lookup["sony"] = 50
	lookup["numa"] = 25
	lookup["kala"] = 15

	power, exists := lookup["foysal"]
	delete(lookup, "kala1")
	fmt.Println(power, exists, len(lookup))

	lookup1 := make(map[int]string, 2) // set initial size
	lookup1[1] = "Foysal"
	lookup1[2] = "Sony"
	lookup1[3] = "Numa" // auto increment
	fmt.Println(lookup1)

	goku := &Saiyan{
		Name:    "Foysal",
		Friends: make(map[string]*Saiyan),
	}

	numa := &Saiyan{
		Name:    "Numa",
		Friends: make(map[string]*Saiyan),
	}

	goku.Friends["Numa"] = numa
	fmt.Println(numa)
	fmt.Println(goku, goku.Friends["Numa"])

	lookup2 := map[string]int{
		"Foysal": 1,
		"Sony":   2,
	}
	fmt.Println(lookup2)

	for key, val := range lookup1 {
		fmt.Println(key, val)
	}

}
