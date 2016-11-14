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
	for i := 0; i < 100; i++ {

		c := big.NewInt(0)
		c.SetString(strconv.Itoa(i), 10)
		f1 := BySpaceOptimize(i).String()
		i1, _ := strconv.Atoi(string(f1[len(f1)-1]))

		f2 := LastDigitOfFibByDinamic(i)

		if i1 != f2 {
			t.Errorf("Wrong asn for %d; Main Fic: %d; last Digit: %d\n", i, f1, f2)
		}

	}
}

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
