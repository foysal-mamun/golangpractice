package maximizing_revenue_in_online_ad_lacement

import "testing"

func TestNativeSolution(t *testing.T) {
	testCases := []struct {
		profitPerClick        []int
		averageNumberOfClicks []int
		expected              int
	}{
		{[]int{23}, []int{39}, 897},
		{[]int{1, 3, -5}, []int{-2, 4, 1}, 23},
		{[]int{1, 2, 3, 4, 5}, []int{1, 0, 1, 0, 1}, 12},
	}

	for _, tc := range testCases {
		actual := NativeSolution(tc.profitPerClick, tc.averageNumberOfClicks)
		if actual != tc.expected {
			t.Errorf("NativeSolution(%v, %v): excepted %d, actual %d", tc.profitPerClick, tc.averageNumberOfClicks, tc.expected, actual)
		}
	}
}
