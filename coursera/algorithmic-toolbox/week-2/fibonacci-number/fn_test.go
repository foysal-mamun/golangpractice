package fibonacci_number

import (
	"math/big"
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
