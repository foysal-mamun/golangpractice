package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	scores := []int{6, 5, 4}
	for i, v := range scores {
		fmt.Printf("scores[%d] = %d\n", i, v)
	}

	scores1 := make([]int, 0, 5)
	c := cap(scores1)
	//scores1 = append(scores1, 5)
	//fmt.Println(scores1)

	for i := 0; i < 25; i++ {
		scores1 = append(scores1, i)

		c1 := cap(scores1)
		if c1 != c {
			c = c1
			fmt.Println(c)
		}
	}

	fmt.Println(scores1)

	slice := scores1[2:4]
	slice[0] = 999
	fmt.Println(scores1)

	source := []int{1, 2, 3, 4, 5}
	source = removeAtIndex(source, 2)
	fmt.Println(source)

	scores2 := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores2[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores2)
	fmt.Println(scores2)

	worst := make([]int, 5)
	copy(worst, scores2[:6])
	fmt.Println(worst)
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
