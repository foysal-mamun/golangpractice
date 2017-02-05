package maximizing_the_number_of_prize

import "math"

func mtnop(n int) (k int, pd []int) {

	p := []int{}
	limit := int(math.Sqrt(float64(n)))
	i := 1
	for ; i <= limit; i = i + 1 {

		if n-i <= i {
			break
		}
		n = n - 1
		p = append(p, i)

	}

	p = append(p, n)

	return i, p
}
