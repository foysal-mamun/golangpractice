package changing_money

import "testing"

func TestNativeSolution(t *testing.T) {
	var testCases = []struct {
		input    int
		expected int
	}{
		{2, 2},
		{28, 6},
		{15684, 1572},
	}

	for _, tc := range testCases {
		actual := NativeSolution(tc.input, 0)
		if actual != tc.expected {
			t.Errorf("NativeSolution(%d): expected %d, actual %d", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkNativeSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NativeSolution(15684, 0)
	}
}
