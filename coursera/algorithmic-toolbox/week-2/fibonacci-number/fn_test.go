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
		t.Errorf("Worng result %v, Correct result %v", f, c)
	}

}

func TestBigIntRecursion(t *testing.T) {

	// c := big.NewInt(354224848179261915075)
	c := big.NewInt(55)
	f := ByRecursionBig(20)
	if c.Cmp(f) != 0 {
		t.Errorf("Worng result %d, Correct result %d", f, c)
	}

	j := big.NewInt(0)
	j.SetString("354224848179261915075", 10) // octal

}
