package fibonacci_number

import "math/big"

func ByRecursion(n int64) int64 {

	if n < 2 {
		return n
	}

	return ByRecursion(n-1) + ByRecursion(n-2)
}

func ByDynamicProgramming(n int) *big.Int {

	f := []*big.Int{big.NewInt(0), big.NewInt(1)}

	for i := 2; i <= n; i++ {
		bi := new(big.Int)
		f = append(f, bi.Add(f[i-1], f[i-2]))
	}

	return f[n]
}

func BySpaceOptimize(n int) *big.Int {

	a := big.NewInt(0)
	b := big.NewInt(1)

	if n == 0 {
		return a
	}

	for i := 2; i <= n; i++ {
		bi := new(big.Int)
		a, b = b, bi.Add(a, b)
	}

	return b
}

func LastDigitOfFibByDinamic(n int) int {
	f := []int{0, 1}
	for i := 2; i <= n; i++ {
		f = append(f, (f[i-1]+f[i-2])%10)
	}

	return f[n]
}
