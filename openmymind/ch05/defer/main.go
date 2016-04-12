package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("defer_test.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
}
