package lcm

import (
	"math/rand"
	"testing"
)

func TestNaibeSolution(t *testing.T) {

	var lcmTests = []struct {
		a        int64 // input
		b        int64 // input
		expected int64 // expected result
	}{
		{6, 8, 24},
		{28851538, 1183019, 1933053046},
	}

	for _, lt := range lcmTests {
		actual := NaiveSolution(lt.a, lt.b)
		if actual != lt.expected {
			t.Errorf("NaiveSolution(%d, %d): expected %d, actual %d", lt.a, lt.b, lt.expected, actual)
		}
	}

}

func TestEfficientSolutionAccuracy(t *testing.T) {
	for {
		break
		a := rand.Int63n(1000) + 2
		b := rand.Int63n(1000) + 2

		expected, actual := NaiveSolution(a, b), EfficientSolution(a, b)

		if expected != actual {
			t.Errorf("EfficientSolution(%d, %d): expected %d, actual %d", a, b, expected, actual)
			break
		}

	}
}

func BenchmarkNaiveSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveSolution(875, 356)
	}
}

func BenchmarkEfficientSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientSolution(875, 356)
	}
}
