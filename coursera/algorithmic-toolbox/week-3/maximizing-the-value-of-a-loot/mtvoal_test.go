package maximizing_the_value_of_a_loot

import "testing"

func TestNativeSolution(t *testing.T) {
	var testCases = []struct {
		loots    []loot
		expected int
	}{
		{[]loot{{60, 20}, {100, 50}, {120, 30}}, 180},
	}

	for _, tc := range testCases {
		actual := NativeSolution(tc.loots)
		if actual != tc.expected {
			t.Errorf("NativeSolution(%v): excepted=%d, actual=%d", tc.loots, tc.expected, actual)
		}
	}

}
