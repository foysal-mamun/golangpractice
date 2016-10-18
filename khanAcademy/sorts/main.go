package main

import (
	"fmt"

	"github.com/foysal-mamun/golangpractice/khanAcademy/sorts/selectionSort"
)

func main() {

	array := []int{22, 11, 99, 88, 9, 7, 42}

	ss := selectionSort.NewSelectionSort(array)
	ss.DoSort()
	fmt.Println(ss.Array)
}
