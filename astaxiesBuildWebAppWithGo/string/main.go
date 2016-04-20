package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) + " Foysal"
	s2 = "Hi " + s2[6:]
	fmt.Println(s2)

}
