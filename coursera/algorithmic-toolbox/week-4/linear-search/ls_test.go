package linear_search

import "testing"

func TestLs(t *testing.T) {
	testCases := []struct {
		n        []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1}, 1, 2},
		{[]int{100, 524, 12, 65}, 65, 3},
	}

	for _, tc := range testCases {
		actual := ls(tc.n, tc.k)
		if actual != tc.expected {
			t.Errorf("ls(%v, %d): expected: %d, actual: %d", tc.n, tc.k, tc.expected, actual)
		}
	}
}
