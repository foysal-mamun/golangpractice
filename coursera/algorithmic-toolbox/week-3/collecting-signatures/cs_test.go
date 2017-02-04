package collecting_signatures

import "testing"

func TestCs(t *testing.T) {
	var testCase = []struct {
		schedules     []schedule
		exceptedCount int
		schedule      schedule
	}{
		{[]schedule{{1, 3}, {2, 5}, {3, 6}}, 1, schedule{3, 0}},
		{[]schedule{{4, 7}, {1, 3}, {2, 5}, {5, 6}}, 2, schedule{3, 6}},
	}

	for _, tc := range testCase {
		actualCount, actualSchedule := cs(tc.schedules)
		if actualCount != tc.exceptedCount || actualSchedule != tc.schedule {
			t.Errorf("cs(%v): excepted: %d and %v but got %d and %v", tc.schedules, tc.exceptedCount, tc.schedule, actualCount, actualSchedule)
		}
	}
}
