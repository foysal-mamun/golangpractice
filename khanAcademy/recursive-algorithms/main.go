package main

import (
	"fmt"

	"github.com/foysal-mamun/golangpractice/khanAcademy/recursive-algorithms/factorial"
)

func main() {
	n := factorial.IterativeFactorial(5)
	fmt.Println("Factorial (Iterative) of 5 is: ", n)

	n = factorial.RecursiveFactorial(5)
	fmt.Println("Factorial (Recursive) of 5 is: ", n)

	p := factorial.NewPalindrome()
	p.Str = "foyof"
	fmt.Printf("foyof is a Palindrome: %v \n", p.IsPalindrome())

	pw := factorial.NewPower()
	fmt.Println("2 Power of -3 is:", pw.GetPower(2, -3))
}
