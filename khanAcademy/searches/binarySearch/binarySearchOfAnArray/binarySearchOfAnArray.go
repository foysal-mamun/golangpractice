package binarySearchOfAnArray

type BinarySearch struct {
	array []int
}

func NewBinarySearch(arr []int) *BinarySearch {
	return &BinarySearch{array: arr}
}

func (b *BinarySearch) DoSearch(target int) int {

	min := 0
	max := len(b.array) - 1

	//return min < max || min == max

	for min < max || min == max {

		guess := int((min + max) / 2)
		guessVal := b.array[guess]

		if guessVal == target {
			return guess
		} else if guessVal < target {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}

	return -1
}
