package fibonacci_number

import (
	"math/big"
	"strconv"
	"testing"
)

func TestByRecursion(t *testing.T) {

	// c := big.NewInt(354224848179261915075)
	var c int64 = 6765
	f := ByRecursion(20)
	if c != f {
		t.Errorf("Wrong result %v, Correct result %v", f, c)
	}

}

func TestByDynamicProgramming(t *testing.T) {

	c := big.NewInt(0)
	c.SetString("354224848179261915075", 10)
	f := ByDynamicProgramming(100)
	if c.Cmp(f) != 0 {
		t.Errorf("Wrong result %d, Correct result %d", f, c)
	}

}

func TestBySpaceOptimize(t *testing.T) {

	c := big.NewInt(0)
	c.SetString("354224848179261915075", 10)
	f := BySpaceOptimize(100)
	if c.Cmp(f) != 0 {
		t.Errorf("Wrong resutl %d, Correct result %d", f, c)
	}
}

func TestLastDigitOfFibByDinamic(t *testing.T) {

	// // Single number checking
	// i := 331
	// f1 := 9

	// f2 := LastDigitOfFibByDinamic(i)

	// if f1 != f2 {
	// 	t.Errorf("Wrong asn for %d; Main Fic: %d; last Digit: %d\n", i, f1, f2)
	// }

	// // checking in loop
	for i := 0; i < 100; i++ {

		c := big.NewInt(0)
		c.SetString(strconv.Itoa(i), 10)
		f1 := BySpaceOptimize(i).String()
		i1, _ := strconv.Atoi(string(f1[len(f1)-1]))

		f2 := LastDigitOfFibByDinamic(i)

		if i1 != f2 {
			t.Errorf("Wrong asn for %d; Main Fic: %s; last Digit: %d\n", i, f1, f2)
		}

	}
}

func TestRepeatedPeriod(t *testing.T) {
	var testCases = []struct {
		m        int
		expected int
	}{
		{2, 3},
		{3, 8},
		{4, 6},
		{5, 20},
		{10, 60},
		{100, 300},
		{1000, 1500},
		{10000, 15000},
	}

	for _, tc := range testCases {
		actual := RepeatedPeriod(tc.m)
		if actual != tc.expected {
			t.Errorf("RepeatedPeriod(%d): expected = %d, actual = %d", tc.m, tc.expected, actual)
		}
	}
}

func TestLastDigitOfFibByPeriod(t *testing.T) {

	// // checking in loop
	for i := 2; i < 310; i++ {

		c := big.NewInt(0)
		c.SetString(strconv.Itoa(i), 10)
		f1 := BySpaceOptimize(i).String()
		expected, _ := strconv.Atoi(string(f1[len(f1)-1]))

		actual := LastDigitOfFibByPeriod(i)

		if expected != actual {
			t.Errorf("f(%d)=%s, LastDigitOfFibByPeriod(%d): expected = %d, actual = %d", i, f1, i, expected, actual)
		}

	}
}

func TestHugeFibModuloNaive(t *testing.T) {

	var hfmnTests = []struct {
		n        int
		m        int64
		expected *big.Int
	}{
		{239, 1000, big.NewInt(161)},
		// {2816213588, 30524, big.NewInt(10249)},
		{1, 239, big.NewInt(1)},
	}

	for _, ht := range hfmnTests {
		actual := HugeFibModuloNaive(ht.n, ht.m)
		if actual.Cmp(ht.expected) != 0 {
			t.Errorf("HugeFibModuloNaive(%d, %d), actual %d, excepted %d", ht.n, ht.m, actual, ht.expected)
		}
	}
}

func TestHugeFibModuloEffecient(t *testing.T) {

	var hfmnTests = []struct {
		n        int
		m        int
		expected int
	}{
		{239, 1000, 161},
		{2816213588, 30524, 10249},
		{1, 239, 1},
	}

	for _, ht := range hfmnTests {
		actual := HugeFibModuloEffecient(ht.n, ht.m)
		if actual != ht.expected {
			t.Errorf("HugeFibModuloEffecient(%d, %d), actual %d, excepted %d", ht.n, ht.m, actual, ht.expected)
		}
	}
}

func TestFibSumLastDigitNaive(t *testing.T) {
	var testCases = []struct {
		n        int
		expected int64
	}{
		{3, 4},
		{100, 5},
		{44, 2},
		{60, 0},
	}

	for _, tc := range testCases {
		actual := FibSumLastDigitNaive(tc.n)
		if actual != tc.expected {
			t.Errorf("FibSumLastDigitNaive(%d): actual %d, expected %d", tc.n, actual, tc.expected)
		}
	}
}

func TestFibSumLastDigitImprovement(t *testing.T) {

	for i := 2; i <= 200; i++ {
		expected := FibSumLastDigitNaive(i)
		actual := FibSumLastDigitImprovement(i)
		if int64(actual) != expected {
			t.Errorf("FibSumLastDigitImprovemetn(%d): actual %d, expected %d", i, actual, expected)
		}
	}
}

func TestFibSumLastDigitMoreImprovement(t *testing.T) {

	for i := 2; i <= 200; i++ {
		expected := FibSumLastDigitImprovement(i)
		actual := FibSumLastDigitMoreImprovement(i)
		if actual != expected {
			t.Errorf("FibSumLastDigitMoreImprovement(%d): actual %d, expected %d", i, actual, expected)
		}
	}

	// var testCases = []struct {
	// 	n        int
	// 	expected int
	// }{
	// 	{3, 4},
	// 	{100, 5},
	// 	{121, 1},
	// 	{224, 2},
	// }

	// for _, tc := range testCases {
	// 	actual := FibSumLastDigitMoreImprovement(tc.n)
	// 	if actual != tc.expected {
	// 		t.Errorf("FibSumLastDigitMoreImprovement(%d): actual %d, expected %d", tc.n, actual, tc.expected)
	// 	}
	// }

	// FibSumLastDigitNaive(121)
	// FibSumLastDigitImprovement(224)
}

func TestLastDigitOfFibPartialSumNaive(t *testing.T) {
	var testCases = []struct {
		s        int
		e        int
		expected int64
	}{
		{3, 7, 1},
		{10, 10, 5},
		{10, 200, 2},
	}

	for _, tc := range testCases {
		actual := LastDigitOfFibPartialSumNaive(tc.s, tc.e)
		if tc.expected != actual {
			t.Errorf("LastDigitOfFibPartialSumNaive(%d, %d): actual %d, expected %d", tc.s, tc.e, actual, tc.expected)
		}
	}
}

func TestLastDigitOfFibPartialSumEfficient(t *testing.T) {
	var testCases = []struct {
		s        int
		e        int
		expected int64
	}{
		{3, 7, 1},
		{10, 10, 5},
		{10, 200, 2},
	}

	for _, tc := range testCases {
		actual := LastDigitOfFibPartialSumEfficient(tc.s, tc.e)
		if tc.expected != int64(actual) {
			t.Errorf("LastDigitOfFibPartialSumEfficient(%d, %d): actual %d, expected %d", tc.s, tc.e, actual, tc.expected)
		}
	}
}

// ----------------- Bench Mark ----------------

func BenchmarkByRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByRecursion(20)
	}

}
func BenchmarkByDynamicProgramming(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ByDynamicProgramming(20)
	}
}

func BenchmarkBySpaceOptimize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BySpaceOptimize(10)
	}
}

func BenchmarkLastDigitOfFibByDinamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastDigitOfFibByDinamic(1000)
	}
}

func BenchmarkLastDigitOfFibByPeriod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastDigitOfFibByPeriod(1000)
	}
}

func BenchmarkHugeFibModuloNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HugeFibModuloNaive(239, 1000)
	}
}

func BenchmarkHugeFibModuloEffecient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HugeFibModuloEffecient(239, 1000)
	}
}

func BenchmarkFibSumLastDigitNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibSumLastDigitNaive(2000)
	}
}

func BenchmarkFibSumLastDigitImprovement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibSumLastDigitImprovement(2000)
	}
}

func BenchmarkFibSumLastDigitMoreImprovement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibSumLastDigitMoreImprovement(2000)
	}
}

func BenchmarkLastDigitOfFibPartialSumNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastDigitOfFibPartialSumNaive(20, 4000)
	}
}

func BenchmarkLastDigitOfFibPartialSumEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastDigitOfFibPartialSumEfficient(20, 4000)
	}
}
