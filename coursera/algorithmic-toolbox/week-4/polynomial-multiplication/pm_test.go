package polynomial_multiplication

import (
	"reflect"
	"testing"
)

func TestNativeSolution(t *testing.T) {
	testCases := []struct {
		A []int
		B []int
		C []int
	}{
		{[]int{5, 0, 10, 6}, []int{1, 2, 4}, []int{5, 10, 30, 26, 52, 24}},
		{[]int{1, 0, 3, 2}, []int{2, 0, 4}, []int{2, 0, 10, 4, 12, 8}},
		{[]int{1, 9, 3, 4, 7}, []int{4, 0, 2, 5}, []int{4, 36, 14, 39, 79, 23, 34, 35}},
		{[]int{4, -5}, []int{2, 3, -6}, []int{8, 2, -39, 30}},
	}

	for _, tc := range testCases {
		c := nativeSolution(tc.A, tc.B)
		if !reflect.DeepEqual(tc.C, c) {
			t.Errorf("nativeSolution(%v, %v): expected %v, actual %v", tc.A, tc.B, tc.C, c)
		}
	}
}
