package fibonacci_number

import "math/big"

func ByRecursion(n int64) int64 {

	if n < 2 {
		return n
	}

	return ByRecursion(n-1) + ByRecursion(n-2)
}

func ByRecursionBig(n int64) *big.Int {

	temp := new(big.Int)

	if n < 2 {
		return big.NewInt(n)
	}

	return temp.Add(ByRecursionBig(n-1), ByRecursionBig(n-2))
}
