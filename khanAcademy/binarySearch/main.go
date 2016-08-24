package main

import (
	"fmt"

	"github.com/foysal-mamun/golangpractice/khanAcademy/binarySearch/binarySearchOfAnArray"
)

func main() {
	fmt.Println("Enter Primary Number for Binary Search: ")
	N := 0
	fmt.Scan(&N)

	array := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	bs := binarySearchOfAnArray.NewBinarySearch(array)
	fmt.Println(bs.DoSearch(N))
}
