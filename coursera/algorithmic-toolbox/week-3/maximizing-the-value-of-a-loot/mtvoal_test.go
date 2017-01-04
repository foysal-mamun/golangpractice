package maximizing_the_value_of_a_loot

import "testing"

func TestNativeSolution(t *testing.T) {
	var testCases = []struct {
		loots    []loot
		capacity int
		expected int
	}{
		{[]loot{{60, 20}, {100, 50}, {120, 30}}, 50, 180},
		{[]loot{{500, 30}}, 10, 166},
	}

	for _, tc := range testCases {
		actual := NativeSolution(tc.loots, tc.capacity)
		if actual != tc.expected {
			t.Errorf("NativeSolution(%v, %d): excepted=%d, actual=%d", tc.loots, tc.capacity, tc.expected, actual)
		}
	}

}
