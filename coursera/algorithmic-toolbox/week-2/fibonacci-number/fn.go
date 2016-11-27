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

/**
 * Pisano period method
 */
func RepeatedPeriod(m int) int {
	a, b, i := 0, 1, 2

	for ; i < m*m; i++ {

		a, b = b, (a+b)%m
		if a == 0 && b == 1 {
			return i - 1
		}
	}

	return i - 1
}

func LastDigitOfFibByPeriod(n int) int {
	m := 10
	period := RepeatedPeriod(m)
	n1 := n % period

	if n1 == 0 {
		return 0
	}

	a, b := 0, 1
	for i := 2; i <= n1; i++ {
		a, b = b, (a+b)%m
	}

	return b
}

func HugeFibModuloNaive(n int, m int64) *big.Int {
	m1 := big.NewInt(m)
	bi := new(big.Int)
	return bi.Mod(BySpaceOptimize(n), m1)
}

func HugeFibModuloEffecient(n, m int) int {
	period := RepeatedPeriod(m)
	n1 := n % period

	if n1 == 0 {
		return 0
	}

	a, b := 0, 1
	for i := 2; i <= n1; i++ {
		a, b = b, (a+b)%m
	}

	return b
}

func FibSumLastDigitNaive(n int) int64 {
	a, b, sum := big.NewInt(0), big.NewInt(1), big.NewInt(1)

	for i := 2; i <= n; i++ {
		a, b = b, a.Add(a, b)
		sum.Add(sum, b)
	}
	return (sum.Mod(sum, big.NewInt(10))).Int64()
}

func FibSumLastDigitImprovement(n int) int {
	a, b, sum := 0, 1, 1

	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%10
		sum += b
		// fmt.Println(i, a, b, sum)
	}

	return sum % 10
}

func FibSumLastDigitMoreImprovement(n int) int {

	m := 10
	rem := n % 60

	if rem == 0 {
		return 0
	}

	a, b, s2 := 0, 1, 1
	for i := 2; i <= rem; i++ {
		a, b = b, (a+b)%m
		s2 += b
	}

	return s2 % m
}

func LastDigitOfFibPartialSumNaive(s, e int) int64 {

	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= s; i++ {
		bi := new(big.Int)
		a, b = b, bi.Add(a, b)
	}

	sum := b
	for i := 0; i < e-s; i++ {
		bi := new(big.Int)
		a, b = b, bi.Add(a, b)

		bi = new(big.Int)
		sum = bi.Add(sum, b)

	}

	return sum.Mod(sum, big.NewInt(10)).Int64()

}

func LastDigitOfFibPartialSumEfficient(s, e int) int {

	period := 60 // as we need last digit, pisano period is 60
	m := 10
	rem := s % period

	a, b := 0, 1
	for i := 2; i <= rem; i++ {
		a, b = b, (a+b)%m
	}

	sum := b
	remE := e % period
	for i := 0; i < remE-rem; i++ {
		a, b = b, (a+b)%m
		sum += b
	}

	return sum % m

}
