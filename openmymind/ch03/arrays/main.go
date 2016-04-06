package main

import (
	"fmt"
)

func main() {

	var scores [5]int
	scores[0] = 1

	fmt.Printf("scores[0] = %d\n", scores[0])

	scores1 := [4]int{9, 8, 7, 6}
	for index, value := range scores1 {
		fmt.Printf("scores[%d] = %d\n", index, value)
	}
}
