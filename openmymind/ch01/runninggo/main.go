package main

import (
	"fmt"
	//"os"
)

func main() {
	println("Bishmillahhir Rahmanirrahim")

	/*if len(os.Args) != 2 {
		fmt.Println(os.Args[0])
		os.Exit(1)
	}

	fmt.Println("over", os.Args[1])*/

	var power int
	power = 9000
	fmt.Printf("This is over %d\n", power)

	var power1 int = 1000
	fmt.Printf("This is over agian %d\n", power1)

	power2 := getPower()
	fmt.Printf("This is over also %d\n", power2)

	name, power2 := "Foysal", getPower()
	fmt.Printf("This is over on name %s, %d\n", name, power2)

	b, a := multiParamReturn(1, "mamun")
	fmt.Printf("Name = %s, Value = %d\n", b, a)
}

func getPower() int {
	return 2000
}

func multiParamReturn(a int, b string) (string, int) {
	return b, a
}
