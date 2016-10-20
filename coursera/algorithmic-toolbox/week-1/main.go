package main

import "fmt"

func main() {

	// A Plus B
	// fmt.Println(aplusb())

	// MaxPairwiseProduct
	N := 0
	fmt.Scan(&N)
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}
	// N := 200000
	// nums := make([]int, N)
	// for i := 0; i < N; i++ {
	// 	nums[i] = i
	// }
	mpp := NewMaxPairwiseProduct(nums)
	fmt.Println(mpp.getMaxPairwiseProduct())
	fmt.Println(mpp.getMaxPairwiseProductFast())

}

func aplusb() int {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	return a + b
}

type MaxPairwiseProduct struct {
	numbers []int
}

func NewMaxPairwiseProduct(n []int) *MaxPairwiseProduct {
	return &MaxPairwiseProduct{numbers: n}
}

func (m *MaxPairwiseProduct) getMaxPairwiseProduct() int {
	product := 0

	numSize := len(m.numbers)
	for i := 0; i < numSize; i++ {
		for j := i + 1; j < numSize; j++ {
			pdct := m.numbers[i] * m.numbers[j]
			if pdct > product {
				product = pdct
			}
		}
	}
	return product
}

func (m *MaxPairwiseProduct) getMaxPair() (int, int) {
	max1, max2 := 0, 0

	numSize := len(m.numbers)
	for i := 0; i < numSize; i++ {

		if m.numbers[i] > max2 {
			max2 = m.numbers[i]
			if max2 > max1 {
				max1, max2 = max2, max1
			}
		}

	}

	return max1, max2
}

func (m *MaxPairwiseProduct) getMaxPairwiseProductFast() int {
	max1, max2 := m.getMaxPair()

	return max1 * max2
}
