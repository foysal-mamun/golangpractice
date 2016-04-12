package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n, err := doubleOne(n1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(n1, n)
	}

	//getInput()
}

func doubleOne(x int) (y int, e error) {

	if x == 1 {
		y = x + x
	} else {
		e = errors.New("Number is not 1")
	}

	return y, e

}

func getInput() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input")
	}
}
