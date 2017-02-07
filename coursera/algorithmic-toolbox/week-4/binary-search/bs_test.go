package binary_search

import "testing"

func TestBs(t *testing.T) {
	testCases := []struct {
		a        []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 3, 4, 4, 5}, 4, 3},
		{[]int{3, 2, 1, 3, 2, 1}, 0, -1},
	}

	for _, tc := range testCases {
		actual := bs(tc.a, tc.k)
		if actual != tc.expected {
			t.Errorf("bs(%v, %d): expected: %d; actual: %d", tc.a, tc.k, tc.expected, actual)
		}
	}
}
