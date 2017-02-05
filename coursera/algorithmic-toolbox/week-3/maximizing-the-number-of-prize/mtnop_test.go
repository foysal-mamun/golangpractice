package maximizing_the_number_of_prize

import (
	"reflect"
	"testing"
)

func TestMtnop(t *testing.T) {
	var testCases = []struct {
		n          int
		exceptedK  int
		exceptedPd []int
	}{
		{6, 3, []int{1, 2, 3}},
		{8, 3, []int{1, 2, 5}},
		{2, 1, []int{2}},
		{1, 1, []int{1}},
		{3, 2, []int{1, 2}},
		{4, 2, []int{1, 3}},
		{5, 2, []int{1, 4}},
	}

	for _, tc := range testCases {
		actualK, actualPd := mtnop(tc.n)
		if actualK != tc.exceptedK || !reflect.DeepEqual(actualPd, tc.exceptedPd) {
			t.Errorf("mtnop(%d): expected K = %d, pd = %v, but actual k = %d, pd = %v", tc.n, tc.exceptedK, tc.exceptedPd, actualK, actualPd)
		}
	}
}
