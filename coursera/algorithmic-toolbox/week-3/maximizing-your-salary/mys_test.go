package maximizing_your_salary

import "testing"

func TestMys(t *testing.T) {
	testCases := []struct {
		nums           []int
		expectedSalary string
	}{
		{[]int{21, 2}, "221"},
		{[]int{9, 4, 6, 1, 9}, "99641"},
		{[]int{23, 39, 92}, "923923"},
	}

	for _, tc := range testCases {
		actualSalary := mys(tc.nums)
		if tc.expectedSalary != actualSalary {
			t.Errorf("myx(%v): expected Salary: %s, actual: %s", tc.nums, tc.expectedSalary, actualSalary)
		}
	}
}
